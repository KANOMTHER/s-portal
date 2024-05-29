package service

import (
	"fmt"
	"log"
	"math/rand"
	"net/http"
	"s-portal/internal/domain/model"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

const SessionKey = "session"
const SessionTokenLen = 32
const SessionTokenAllowedCharacter = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ123456789"

type AuthUser struct {
	UserId uint   `json:"id"`
	Role   string `json:"role"`
}

type AuthService struct {
	Db       *gorm.DB
	Sessions map[string]AuthUser
}

func NewAuthService(db *gorm.DB) *AuthService {
	return &AuthService{
		Db:       db,
		Sessions: make(map[string]AuthUser),
	}
}

func (m *AuthService) GenerateToken() string {
	token := ""
	allowedCharacterLen := len(SessionTokenAllowedCharacter)

	for i := 0; i < SessionTokenLen; i++ {
		index := rand.Int31n(int32(allowedCharacterLen))
		token += string(SessionTokenAllowedCharacter[index])
	}

	return token
}

func (m *AuthService) ValidateUser(userId uint, password string) (*AuthUser, error) {
	var user model.User
	err := m.Db.First(&user, userId).Error
	if err != nil {
		return nil, fmt.Errorf("invalid user %d", userId)
	}

	valid, err := user.Validate(password)

	if err != nil {
		return nil, fmt.Errorf("failed to validate password %w", err)
	}

	if !valid {
		return nil, nil
	}

	return &AuthUser{
		UserId: userId,
		Role:   user.Role,
	}, nil
}

func (m *AuthService) SetContextUser(ctx *gin.Context, user AuthUser) error {
	token := m.GenerateToken()
	host := ctx.Request.Host
	log.Println("Host", host)
	ctx.SetSameSite(http.SameSiteLaxMode)
	ctx.SetCookie(SessionKey, token, 60*60*24*30 /* 30 days */, "/", host, false, true)
	m.Sessions[token] = user
	return nil
}

func (m *AuthService) GetContextUser(ctx *gin.Context) (*AuthUser, error) {
	cookie, err := ctx.Cookie(SessionKey)
	if err != nil {
		return nil, nil
	}

	if len(cookie) != SessionTokenLen {
		return nil, fmt.Errorf("invalid token size, expecting %d but found %d", SessionTokenLen, len(cookie))
	}

	authUser, _ := m.Sessions[cookie]
	return &authUser, nil
}

func (m *AuthService) UnsetContextUser(ctx *gin.Context) error {
	host := ctx.Request.Host
	cookie, err := ctx.Cookie(SessionKey)
	if err != nil {
		return fmt.Errorf("failed to get cookie %w", err)
	}
	ctx.SetCookie(SessionKey, "", -1, "/", host, false, true)
	delete(m.Sessions, cookie)
	return nil
}

func (m *AuthService) AssertPermission(ctx *gin.Context, role ...string) bool {
	user, err := m.GetContextUser(ctx)
	if err != nil {
		log.Println("failed asserting role", err)
		return false
	}

	if user == nil {
		return false
	}

	if len(role) == 0 {
		return true
	}

	for _, s := range role {
		if user.Role == s {
			return true
		}
	}
	return false
}
