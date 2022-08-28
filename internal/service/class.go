package service

import (
	"errors"
	"github.com/google/uuid"
	"gorm.io/gorm"
	"main/internal/entity"
)

type ClassService struct {
	Db *gorm.DB
}

func (s *ClassService) Create(class entity.Class) error {
	if err := s.Db.Create(&class); err.Error != nil {
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

	if err := s.Db.First(&class, "id = ?", uuid); err.Error != nil {
		return entity.Class{}, errors.New("error in get by id on database")
	}

	return class, nil
}

func (s *ClassService) GetAll() ([]entity.Class, error) {
	var classes []entity.Class

	if err := s.Db.Find(&classes); err.Error != nil {
		return []entity.Class{}, errors.New("error in get by id on database")
	}

	return classes, nil
}

func (s *ClassService) Delete(class entity.Class) error {
	class.Deleted = true
	if err := s.Update(class); err != nil {
		return err
	}

	return nil
}
