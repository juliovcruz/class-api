package response

import (
	"github.com/google/uuid"
	"main/internal/entity"
	"time"
)

type ClassResponse struct {
	ID          uuid.UUID `json:"id"`
	Name        string    `json:"name"`
	TotalHours  int       `json:"total_hours"`
	Schedules   string    `json:"schedules"` // 2N2345
	Description string    `json:"description"`
	TeacherID   uuid.UUID `json:"teacher_id"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}

func NewClassResponse(class entity.Class) ClassResponse {
	return ClassResponse{
		ID:          class.ID,
		Name:        class.Name,
		TotalHours:  class.TotalHours,
		Description: class.Description,
		Schedules:   class.Schedules,
		TeacherID:   class.TeacherID,
		CreatedAt:   class.CreatedAt,
		UpdatedAt:   class.UpdatedAt,
	}
}

func NewClassesResponse(classes []entity.Class) []ClassResponse {
	res := []ClassResponse{}

	for _, class := range classes {
		res = append(res, ClassResponse{
			ID:          class.ID,
			Name:        class.Name,
			TotalHours:  class.TotalHours,
			Description: class.Description,
			Schedules:   class.Schedules,
			TeacherID:   class.TeacherID,
			CreatedAt:   class.CreatedAt,
			UpdatedAt:   class.UpdatedAt,
		})
	}

	return res
}
