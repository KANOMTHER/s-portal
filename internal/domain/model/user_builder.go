package model

import "strconv"

// Builder Interface
type UserBuilder interface {
	setID(id uint)
	setPWD()
	setRole()
	getUser() User
}

func GetUserBuilder(builderType string) UserBuilder {
	if builderType == "admin" {
		return newAdminBuilder()
	}

	if builderType == "student" {
		return newStudentBuilder()
	}

	if builderType == "teacher" {
		return newTeacherBuilder()
	}

	return nil
}

// Concrete Builder : AdminBuilder
type AdminBuilder struct {
	user User
}

func newAdminBuilder() *AdminBuilder {
	return &AdminBuilder{}
}

func (a *AdminBuilder) setID(id uint) {
	a.user.ID = id
}

func (a *AdminBuilder) setPWD() {
	// Admin's password is the same as the ID
	authDecorator := AuthDecorator{User: &a.user}
	authDecorator.SetPassword(strconv.FormatUint(uint64(a.user.ID), 10))
}

func (a *AdminBuilder) setRole() {
	a.user.Role = "admin"
}

func (a *AdminBuilder) getUser() User {
	return User{
		ID:   a.user.ID,
		PWD:  a.user.PWD,
		Role: a.user.Role,
	}
}

// Concrete Builder : StudentBuilder
type StudentBuilder struct {
	user User
}

func newStudentBuilder() *StudentBuilder {
	return &StudentBuilder{}
}

func (s *StudentBuilder) setID(id uint) {
	s.user.ID = id
}

func (s *StudentBuilder) setPWD() {
	// Student's password is the same as the ID
	authDecorator := AuthDecorator{User: &s.user}
	authDecorator.SetPassword(strconv.FormatUint(uint64(s.user.ID), 10))
}

func (s *StudentBuilder) setRole() {
	s.user.Role = "student"
}

func (s *StudentBuilder) getUser() User {
	return User{
		ID:   s.user.ID,
		PWD:  s.user.PWD,
		Role: s.user.Role,
	}
}

// Concrete Builder : TeacherBuilder
type TeacherBuilder struct {
	user User
}

func newTeacherBuilder() *TeacherBuilder {
	return &TeacherBuilder{}
}
func (t *TeacherBuilder) setID(id uint) {
	t.user.ID = id
}
func (t *TeacherBuilder) setPWD() {
	// Teacher's password is the same as the ID
	authDecorator := AuthDecorator{User: &t.user}
	authDecorator.SetPassword(strconv.FormatUint(uint64(t.user.ID), 10))
}
func (t *TeacherBuilder) setRole() {
	t.user.Role = "teacher"
}
func (t *TeacherBuilder) getUser() User {
	return User{
		ID:   t.user.ID,
		PWD:  t.user.PWD,
		Role: t.user.Role,
	}
}
