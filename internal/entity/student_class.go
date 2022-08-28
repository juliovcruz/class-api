package entity

import (
	"database/sql"
	"github.com/google/uuid"
	"time"
)

type StudentClass struct {
	ID        uuid.UUID    `gorm:"primaryKey;type:uuid"`
	ClassID   uuid.UUID    `gorm:"column:class_id"`
	StudentID uuid.UUID    `gorm:"column:student_id"`
	CreatedAt time.Time    `gorm:"column:created_at"`
	UpdatedAt time.Time    `gorm:"column:updated_at"`
	DeletedAt sql.NullTime `gorm:"column:deleted_at"`
}
