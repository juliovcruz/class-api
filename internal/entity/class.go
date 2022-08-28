package entity

import (
	"database/sql"
	"github.com/google/uuid"
	"time"
)

type Class struct {
	ID          uuid.UUID `gorm:"primaryKey;type:uuid"`
	Name        string
	Description string
	TotalHours  int    `gorm:"column:total_hours"`
	Schedules   string // 2N2345
	TeacherID   uuid.UUID
	CreatedAt   time.Time    `gorm:"column:created_at"`
	UpdatedAt   time.Time    `gorm:"column:updated_at"`
	DeletedAt   sql.NullTime `gorm:"column:deleted_at"`
}
