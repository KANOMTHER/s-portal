package model

import "github.com/alexedwards/argon2id"

type AuthDecorator struct {
	User *User
}

func (a *AuthDecorator) Validate(password string) (bool, error) {
	hash := a.User.PWD

	return argon2id.ComparePasswordAndHash(password, hash)
}

func (a *AuthDecorator) SetPassword(password string) error {
	hash, err := argon2id.CreateHash(password, argon2id.DefaultParams)
	if err != nil {
		return err
	}

	a.User.PWD = hash
	return nil
}
