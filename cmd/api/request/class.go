package request

import (
	"database/sql"
	"github.com/google/uuid"
	"main/internal/entity"
	"time"
)

type ClassRequest struct {
	Name        string `validate:"required,gte=5,lte=50"`
	Description string `validate:"required,gte=5,lte=255"` // TODO: DB
	TotalHours  int    `validate:"required" json:"total_hours"`
	Schedules   string
	TeacherID   string `json:"teacher_id"`
}

func (r *ClassRequest) ToEntity() (entity.Class, error) {
	id, err := uuid.Parse(r.TeacherID)
	if err != nil {
		return entity.Class{}, err
	}

	return entity.Class{
		ID:          uuid.New(),
		Name:        r.Name,
		TotalHours:  r.TotalHours,
		Description: r.Description,
		Schedules:   r.Schedules,
		TeacherID:   id,
		DeletedAt:   sql.NullTime{},
		CreatedAt:   time.Now(),
		UpdatedAt:   time.Now(),
	}, nil
}
