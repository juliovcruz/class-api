package service

import (
	"database/sql"
	"errors"
	"github.com/google/uuid"
	"gorm.io/gorm"
	"main/internal/entity"
	"strings"
	"time"
)

type ClassService struct {
	Db *gorm.DB
}

func (s *ClassService) Create(class entity.Class) error {
	if err := s.Db.Create(&class); err.Error != nil {
		if strings.Contains(err.Error.Error(), "duplicate") {
			return errors.New("class already exists")
		}

		return errors.New("error in insert class on database")
	}

	return nil
}

func (s *ClassService) Update(class entity.Class) error {
	if err := s.Db.Save(&class); err.Error != nil {
		return errors.New("error in update class on database")
	}

	return nil
}

func (s *ClassService) GetByID(uuid uuid.UUID) (entity.Class, error) {
	var class entity.Class

	if err := s.Db.First(&class, "id = ? AND deleted_at is null", uuid); err.Error != nil {
		if errors.Is(err.Error, gorm.ErrRecordNotFound) {
			return entity.Class{}, err.Error
		}

		return entity.Class{}, errors.New("error in get by id on database")
	}

	return class, nil
}

func (s *ClassService) GetAll() ([]entity.Class, error) {
	var classes []entity.Class

	if err := s.Db.Find(&classes, "deleted_at is null"); err.Error != nil {
		return []entity.Class{}, errors.New("error in get by id on database")
	}

	return classes, nil
}

func (s *ClassService) Delete(class entity.Class) error {
	class.DeletedAt = sql.NullTime{
		Time:  time.Now(),
		Valid: true,
	}

	if err := s.Update(class); err != nil {
		return err
	}

	return nil
}
