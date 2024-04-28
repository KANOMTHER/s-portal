package model

type UserDirector struct {
	builder UserBuilder
}

func NewUserDirector(builder UserBuilder) *UserDirector {
	return &UserDirector{
		builder: builder,
	}
}

func (director *UserDirector) SetUserBuilder(builder UserBuilder) {
	director.builder = builder
}

func (director *UserDirector) Construct(id uint) User {
	director.builder.setID(id)
	director.builder.setPWD()
	director.builder.setRole()
	return director.builder.getUser()
}