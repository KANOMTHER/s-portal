package service

import (
	"fmt"
	"net/url"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"

	"s-portal/internal/domain/model"
)

type TimeTableService struct {
	db *gorm.DB
}

func NewTimeTableService(db *gorm.DB) *TimeTableService {
	return &TimeTableService{
		db: db,
	}
}

type GetTimetableByClassIDField struct {
	CourseCode string       `example:"CPE313"`
	Section    string       `example:"A"`
	Day        time.Weekday `example:"0"`
	StartTime  time.Time    `example:"2021-08-01T08:00:00Z"`
	EndTime    time.Time    `example:"2021-08-01T09:00:00Z"`
	Classroom  string       `example:"CPE1102"`
	ClassType  string       `example:"Lecture"`
}

func (ts *TimeTableService) GetTimetableByClassID(class_id string) ([]GetTimetableByClassIDField, error) {
	var timetables []GetTimetableByClassIDField
	if err := ts.db.
		Model(&model.Timetable{}).
		Select("courses.course_code AS CourseCode, classes.section AS Section, timetables.day AS Day, timetables.start_time AS StartTime, timetables.end_time AS EndTime, timetables.classroom AS Classroom, timetables.class_type AS ClassType").
		Joins("inner join classes on classes.id = timetables.class_id").
		Joins("inner join courses on courses.id = classes.course_id").
		Where("timetables.class_id = ?", class_id).
		Order("CASE Day WHEN 0 THEN 7 ELSE Day END ASC, StartTime ASC").
		Scan(&timetables).Error; err != nil {
		return nil, err
	}

	if timetables == nil {
		var not_found = make([]GetTimetableByClassIDField, 0)
		return not_found, nil
	}

	return timetables, nil
}

func (ts *TimeTableService) GetStudentTimetable(context *gin.Context) ([]GetTimetableByClassIDField, error) {
	ps := NewClassRegisterService(ts.db)

	register_classes, retErr := ps.GetRegisterClassByID(context)
	if retErr != nil {
		return nil, retErr
	}

	student_timeTable := []GetTimetableByClassIDField{}

	for i := 0; i < len(register_classes); i++ {
		class_timetable, retErr := ts.GetTimetableByClassID(strconv.FormatUint(uint64(register_classes[i].ClassID), 10))
		if retErr != nil {
			return nil, retErr
		}

		for j := 0; j < len(class_timetable); j++ {
			student_timeTable = append(student_timeTable, class_timetable[j])
		}
	}

	return student_timeTable, nil
}

func (ts *TimeTableService) GetTATimetable(context *gin.Context) ([]GetTimetableByClassIDField ,error) {
	payment := model.Payment{}

	if err := context.ShouldBindJSON(&payment); err != nil {
        return nil, err
    }

	existPaymentID := model.Payment{}
	if err := ts.db.Model(model.Payment{}).Where(&payment).First(&existPaymentID).Error; err != nil {
		return nil, err
    }

	// get class id
	var classID []uint64
 	if err := ts.db.Model(model.ClassRegister{}).Where("payment_id = ?", existPaymentID.ID).Select("id").Find(&classID).Error ; err!=nil {
		return nil, err
	}

	// get time
	ta_timeTable := []GetTimetableByClassIDField{}

	for i := 0; i < len(classID); i++ {
		class_timetable, retErr := ts.GetTimetableByClassID(strconv.FormatUint(classID[i], 10))
		if retErr != nil {
			return nil, retErr
		}

		for j := 0; j < len(class_timetable); j++ {
			ta_timeTable = append(ta_timeTable, class_timetable[j])
		}
	}

	return ta_timeTable, nil
}

func (ts *TimeTableService) CreateTimeTable(timeTable *model.Timetable) error {
	if err := ts.db.Create(&timeTable).Error; err != nil {
		return err
	}
	return nil
}

func (ts *TimeTableService) CountTimeTable(queryParams url.Values) (int64, error) {
	var count int64
	query := ts.db.Model(&model.Timetable{})
	for key, values := range queryParams {
		for _, value := range values {
			query = query.Where(key+" = ?", value)
		}
	}
	if err := query.Count(&count).Error; err != nil {
		return -1, err
	}
	return count, nil
}

func (ts *TimeTableService) DeleteTimeTableByID(id string) error {
	timeTable := model.Timetable{}
	if result := ts.db.Delete(&timeTable, id); result.RowsAffected < 1 {
		return fmt.Errorf("were not able to delete the timeTable")
	}

	return nil
}
