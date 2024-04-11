package service

import (
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

func (ts *TimeTableService) CreateTimeTable(timeTable *model.Timetable) error {
	if err := ts.db.Create(&timeTable).Error; err != nil {
		return err
	}
	return nil
}

func (ts *TimeTableService) CountTimeTable(queryParams url.Values) (int64, error) {
	var count int64
	query := ts.db.Model(&model.Timetable{}) // Replace YourModel with your GORM model struct
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