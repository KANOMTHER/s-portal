package model

import (
	"github.com/alexedwards/argon2id"
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	ID   uint `gorm:"primaryKey"`
	PWD  string
	Role string
}

func (u *User) Validate(password string) (bool, error) {
	hash := u.PWD

	return argon2id.ComparePasswordAndHash(password, hash)
}

func (u *User) SetPassword(password string) error {
	hash, err := argon2id.CreateHash(password, argon2id.DefaultParams)
	if err != nil {
		return err
	}

	u.PWD = hash
	return nil
}
