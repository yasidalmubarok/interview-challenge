package service

import (
	"challenge-interview/dto/request"
	"challenge-interview/dto/response"
	"challenge-interview/entity"
	"challenge-interview/helper"
	"challenge-interview/repository"
	"time"
)

type OrderService interface {
	Create(payload *request.OrderRequest) (*response.OrderResponse, helper.Error)
	GetByID(id int) (*response.OrderResponse, helper.Error)
	GetAll() ([]*response.OrderResponse, helper.Error)
	Update(payload *request.OrderRequest, id int) (*response.OrderResponse, helper.Error)
	Delete(id int) (bool, helper.Error)
}

type orderService struct {
	orderRepo repository.OrderRepository
	carRepo   repository.CarRepository
}

func NewOrderService(orderRepo repository.OrderRepository, carRepo repository.CarRepository) *orderService {
	return &orderService{
		orderRepo: orderRepo,
		carRepo:   carRepo,
	}
}

func (s *orderService) Create(payload *request.OrderRequest) (*response.OrderResponse, helper.Error) {
	order := &entity.Order{}
	order.CarID = payload.CarID
	order.DropoffDate = payload.DropoffDate
	order.PickupDate = payload.PickupDate
	order.PickupLocation = payload.PickupLocation
	order.DropoffLocation = payload.DropoffLocation
	order.OrderDate = time.Now()
	order.CreatedAt = time.Now()
	order.UpdatedAt = time.Now()

	if order.PickupDate.After(order.DropoffDate) {
		return nil, helper.NewStatusUnProcessableEntityError("Pickup date cannot be after dropoff date")
	}

	_, err := s.carRepo.GetByID(order.CarID)
	if err != nil {
		return nil, helper.NewStatusNotFoundError("Car not found")
	}

	createdOrder, err := s.orderRepo.Create(order)
	if err != nil {
		return nil, err
	}

	return response.MapOrderToResponse(createdOrder), nil
}

func (s *orderService) GetByID(id int) (*response.OrderResponse, helper.Error) {
	order, err := s.orderRepo.GetByID(id)
	if err != nil {
		return nil, err
	}

	return response.MapOrderToResponse(order), nil
}

func (s *orderService) GetAll() ([]*response.OrderResponse, helper.Error) {
	orders, err := s.orderRepo.GetAll()
	if err != nil {
		return nil, err
	}

	var orderResponses []*response.OrderResponse
	for _, order := range orders {
		orderResponses = append(orderResponses, response.MapOrderToResponse(order))
	}

	return orderResponses, nil
}

func (s *orderService) Update(payload *request.OrderRequest, id int) (*response.OrderResponse, helper.Error) {
	order, err := s.orderRepo.GetByID(id)
	if err != nil {
		return nil, err
	}

	order.CarID = payload.CarID
	order.PickupDate = payload.PickupDate
	order.DropoffDate = payload.DropoffDate
	order.PickupLocation = payload.PickupLocation
	order.DropoffLocation = payload.DropoffLocation
	_, err = s.carRepo.GetByID(order.CarID)
	if err != nil {
		return nil, helper.NewStatusNotFoundError("Car not found")
	}

	updatedOrder, err := s.orderRepo.Update(order, id)
	if err != nil {
		return nil, err
	}

	return response.MapOrderToResponse(updatedOrder), nil
}

func (s *orderService) Delete(id int) (bool, helper.Error) {
	_, err := s.orderRepo.GetByID(id)
	if err != nil {
		return false, err
	}

	deleted, err := s.orderRepo.Delete(id)
	if err != nil {
		return false, err
	}
	if !deleted {
		return false, helper.NewStatusNotFoundError("order not found")
	}
	return true, nil
}
