package service

import (
	"challenge-interview/dto/request"
	"challenge-interview/dto/response"
	"challenge-interview/entity"
	"challenge-interview/helper"
	"challenge-interview/repository"
	"time"
)

type CarService interface {
	Create(payload *request.CarRequest) (*response.CarResponse, helper.Error)
	GetAll() ([]*response.CarResponse, helper.Error)
	GetByID(id int) (*response.CarResponse, helper.Error)
	Update(payload *request.CarRequest, id int) (*response.CarResponse, helper.Error)
	Delete(id int) (bool, helper.Error)
}

type carService struct {
	carRepo repository.CarRepository
}

func NewCarService(carRepo repository.CarRepository) *carService {
	return &carService{
		carRepo: carRepo,
	}
}

func (s *carService) Create(payload *request.CarRequest) (*response.CarResponse, helper.Error) {
	newCar := &entity.Car{}
	newCar.Name = payload.Name
	newCar.DayRate = payload.DayRate
	newCar.MonthRate = payload.MonthRate
	newCar.Image = payload.Image
	newCar.CreatedAt = time.Now()
	newCar.UpdatedAt = time.Now()

	createdCar, err := s.carRepo.Create(newCar)
	if err != nil {
		return nil, err
	}

	return response.MapCarToResponse(createdCar), nil
}

func (s *carService) GetAll() ([]*response.CarResponse, helper.Error) {
	cars, err := s.carRepo.GetAll()
	if err != nil {
		return nil, err
	}

	var carResponses []*response.CarResponse
	for _, car := range cars {
		carResponses = append(carResponses, response.MapCarToResponse(car))
	}

	return carResponses, nil
}

func (s *carService) GetByID(id int) (*response.CarResponse, helper.Error) {
	car, err := s.carRepo.GetByID(id)
	if err != nil {
		return nil, err
	}

	return response.MapCarToResponse(car), nil
}

func (s *carService) Update(payload *request.CarRequest, id int) (*response.CarResponse, helper.Error) {
	car, err := s.carRepo.GetByID(id)
	if err != nil {
		return nil, err
	}

	car.Name = payload.Name
	car.DayRate = payload.DayRate
	car.MonthRate = payload.MonthRate
	car.Image = payload.Image
	car.UpdatedAt = time.Now()

	updatedCar, err := s.carRepo.Update(car, id)
	if err != nil {
		return nil, err
	}

	return response.MapCarToResponse(updatedCar), nil
}

func (s *carService) Delete(id int) (bool, helper.Error) {
	_, err := s.carRepo.GetByID(id)
	if err != nil {
		return false, err
	}

	deleted, err := s.carRepo.Delete(id)
	if err != nil {
		return false, err
	}

	return deleted, nil
}