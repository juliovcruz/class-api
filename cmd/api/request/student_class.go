package request

import (
	"database/sql"
	"github.com/google/uuid"
	"main/internal/entity"
	"time"
)

type StudentClassRequest struct {
	ClassID   string `gorm:"column:class_id"`
	StudentID string `gorm:"column:student_id"`
}

func (r *StudentClassRequest) ToEntity() (entity.StudentClass, error) {
	classID, err := uuid.Parse(r.ClassID)
	if err != nil {
		return entity.StudentClass{}, err
	}

	studentID, err := uuid.Parse(r.StudentID)
	if err != nil {
		return entity.StudentClass{}, err
	}

	return entity.StudentClass{
		ID:        uuid.New(),
		ClassID:   classID,
		StudentID: studentID,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
		DeletedAt: sql.NullTime{},
	}, nil
}
