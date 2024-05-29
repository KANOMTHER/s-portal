package model

type GradeInterface interface {
	GradeFromRegisterAdapter() Grade
	GradeStrategy() Grade
}

type GradeHub struct {
	gh GradeInterface
}

func (f *GradeHub) GradeFromRegisterAdapter(classID uint, studentID uint, grade float32) Grade {
	return Grade{
		StudentID: studentID,
		ClassID:   classID,
		Grade:     grade,
	}
}

func (f *GradeHub) GradeStrategy(classRegister ClassRegister, studentID uint, newClass uint) Grade {
	return Grade{
		StudentID: studentID,
		ClassID:   classRegister.ClassID,
		Grade:     0,
	}
}
