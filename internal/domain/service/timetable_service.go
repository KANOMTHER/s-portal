package service

import (
	"time"
	"fmt"
	"net/url"

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
    CourseCode string
    Section   string
    Day       time.Weekday
    StartTime time.Time
    EndTime   time.Time
    Classroom string
    ClassType string
}

func (ts *TimeTableService) GetTimetableByClassID(class_id string) ([]GetTimetableByClassIDField, error) {
	var timetables []GetTimetableByClassIDField
	if err := ts.db.
	Model(&model.Timetable{}).
	Select("courses.course_code AS CourseCode, classes.section AS Section, timetables.day AS Day, timetables.start_time AS StartTime, timetables.end_time AS EndTime, timetables.classroom AS Classroom, timetables.class_type AS ClassType").
	Joins("inner join classes on classes.id = timetables.class_id").
	Joins("inner join courses on courses.id = classes.course_id").
	Where("timetables.class_id = ?", class_id).
	Scan(&timetables).Error; err != nil {
		return nil, err
	}

	if(timetables == nil){
		return nil, fmt.Errorf("no timetable found for the class")
	}

	return timetables, nil
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