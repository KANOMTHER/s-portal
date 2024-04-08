package model

type StaffFacade interface {
	GetClassID() uint
	GetClass() Class
}

func (t*TA) GetClassID() uint{
	return t.ClassID
}

func (t*TA) GetClass() Class{
	return t.Class
}

func (i*Instructor) GetClassID() uint{
	return i.ClassID
}

func (i*Instructor) GetClass() Class{
	return i.Class
}
