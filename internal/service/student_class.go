package service

import (
	"database/sql"
	"errors"
	"github.com/google/uuid"
	"gorm.io/gorm"
	"main/internal/entity"
	"time"
)

type StudentClassService struct {
	Db *gorm.DB
}

func (s *StudentClassService) Create(studentClass entity.StudentClass) error {
	if err := s.Db.Create(studentClass); err.Error != nil {
		return errors.New("error in insert class on database")
	}

	return nil
}

func (s *StudentClassService) Delete(studentID uuid.UUID, classID uuid.UUID) error {
	var studentClass entity.StudentClass

	if err := s.Db.First(&studentClass, "student_id = ? AND class_id = ?", studentID, classID); err.Error != nil {
		return errors.New("error in get by studentID and classID on database")
	}

	studentClass.DeletedAt = sql.NullTime{
		Time:  time.Now(),
		Valid: true,
	}

	if err := s.Db.Save(studentClass); err.Error != nil {
		return errors.New("error in update class on database")
	}

	return nil
}

func (s *StudentClassService) GetAllByClassID(classID uuid.UUID) ([]entity.StudentClass, error) {
	studentClasses := []entity.StudentClass{}

	if err := s.Db.Find(&studentClasses, "class_id = ?", classID); err.Error != nil {
		return []entity.StudentClass{}, errors.New("error in get by classID on database")
	}

	return studentClasses, nil
}
