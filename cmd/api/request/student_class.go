package request

import (
	"database/sql"
	"github.com/google/uuid"
	"main/internal/entity"
	"time"
)

type StudentClassRequest struct {
	ClassID   string `validate:"uuid4" example:"81136ed8-0df4-4892-8a4a-258a986ec440"`
	StudentID string `validate:"uuid4" example:"8c113d2f-c749-4f29-9e9c-c7d0a4ebc114"`
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
