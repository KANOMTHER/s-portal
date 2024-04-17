package model

type AuthDecorator struct {
	User *User
}

func (a *AuthDecorator) Validate(password string) (bool, error) {
	hash := a.User.PWD

	//TODO add a proper hash check :yay:
	return hash == password, nil
}

func (a *AuthDecorator) SetPassword(password string) error {
	a.User.PWD = password
	return nil
}
