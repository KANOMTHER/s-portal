package model

type StaffBuilder struct {
	classID uint
	class Class
	staffId uint
}

func NewStaffBuilder() StaffBuilder{
	return StaffBuilder{

	}
}

func (b*StaffBuilder) WithClassId(id uint) *StaffBuilder{
	b.classID = id
	return b;
}

func (b*StaffBuilder) WithClass(class Class) *StaffBuilder {
	b.class = class;
	return b;
}

func (b*StaffBuilder) WithStaffId(staffId uint) *StaffBuilder{
	b.staffId = staffId;
	return b;
}

func (b*StaffBuilder) BuildInstructor(professor Professor) Instructor{
	return Instructor{
		ClassID: b.classID,
		Class: b.class,
		ProfessorID: b.staffId,
		Professor: professor,
	}
}

func (b*StaffBuilder) BuildTA(student Student) TA{
	return TA{
		ClassID: b.classID,
		Class: b.class,
		StudentID: b.staffId,
		Student: student,
	}
}
