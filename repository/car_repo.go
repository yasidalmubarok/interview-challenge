package repository

import (
	"challenge-interview/entity"
	"challenge-interview/helper"

	"gorm.io/gorm"
)

type CarRepository interface {
	Create(car *entity.Car) (*entity.Car, helper.Error)
	GetAll() ([]*entity.Car, helper.Error)
	GetByID(id int) (*entity.Car, helper.Error)
	Update(car *entity.Car, id int) (*entity.Car, helper.Error)
	Delete(id int) (bool, helper.Error)
}

type carRepository struct {
	db *gorm.DB
}

func NewCarRepository(db *gorm.DB) *carRepository {
	return &carRepository{
		db: db,
	}
}

func (r *carRepository) Create(car *entity.Car) (*entity.Car, helper.Error) {
	if err := r.db.Create(car).Error; err != nil {
		return nil, helper.NewStatusInternalServerError("Failed to create car: ")
	}
	return car, nil
}

func (r *carRepository) GetAll() ([]*entity.Car, helper.Error) {
	var cars []*entity.Car
	if err := r.db.Find(&cars).Error; err != nil {
		return nil, helper.NewStatusInternalServerError("Failed to retrieve cars: ")
	}
	return cars, nil
}

func (r *carRepository) GetByID(id int) (*entity.Car, helper.Error) {
	var car entity.Car
	if err := r.db.First(&car, id).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, helper.NewStatusNotFoundError("Car not found")
		}
		return nil, helper.NewStatusInternalServerError("Failed to retrieve car: ")
	}
	return &car, nil
}

func (r *carRepository) Update(car *entity.Car, id int) (*entity.Car, helper.Error) {
	if err := r.db.Model(&entity.Car{}).Where("id = ?", id).Updates(car).Error; err != nil {
		return nil, helper.NewStatusInternalServerError("Failed to update car: ")
	}
	return car, nil
}

func (r *carRepository) Delete(id int) (bool, helper.Error) {
	if err := r.db.Delete(&entity.Car{}, id).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return false, helper.NewStatusNotFoundError("Car not found")
		}
		return false, helper.NewStatusInternalServerError("Failed to delete car: ")
	}
	return true, nil
}