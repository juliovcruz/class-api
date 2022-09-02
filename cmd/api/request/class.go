package request

import (
	"database/sql"
	"github.com/google/uuid"
	"main/internal/entity"
	"time"
)

type ClassRequest struct {
	Name        string `validate:"required,gte=5,lte=50" example:"Desenvolvimento Full Stack"`
	Description string `validate:"required,gte=5,lte=255" example:"Ensino de desenvolvimento back-end e front-end"`
	TotalHours  int    `validate:"required" json:"total_hours" example:"123"`
	Schedules   string `example:"2N2345"`
	TeacherID   string `validate:"uuid4" json:"teacher_id" example:"81136ed8-0df4-4892-8a4a-258a986ec440"`
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
