package model

type AuthDecorator struct {
	User *User
}

func (a*AuthDecorator) Validate(password string) bool{
	hash := a.User.PWD
	
	//TODO add a proper hash check :yay:
	return hash == password
}

func (a*AuthDecorator) SetPassword(password string){
	a.User.PWD = password;
}