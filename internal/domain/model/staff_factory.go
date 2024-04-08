package model;

type AbstractStaffFactory interface {
	CreateInstructor() Instructor;
	CreateAssistance() TA;
}

type StaffFactory struct {}

func (f*StaffFactory) CreateInstructor(classID uint, class Class, professorID uint, professor Professor) Instructor{
	return Instructor{
		ClassID: classID,
		Class: class,
		ProfessorID: professorID,
		Professor: professor,
	}
}

func (f*StaffFactory) CreateAssistance(classID uint, class Class, studentId uint, student Student) TA{
	return TA{
		ClassID: classID,
		Class: class,
		StudentID: studentId,
		Student: student,
	}
}