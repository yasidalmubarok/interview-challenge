package repository

import (
	"challenge-interview/entity"
	"challenge-interview/helper"

	"gorm.io/gorm"
)

type OrderRepository interface {
	Create(order *entity.Order) (*entity.Order, helper.Error)
	GetByID(id int) (*entity.Order, helper.Error)
	GetAll() ([]*entity.Order, helper.Error)
	Update(order *entity.Order, id int) (*entity.Order, helper.Error)
	Delete(id int) (bool, helper.Error)
}

type orderRepository struct {
	db *gorm.DB
}

func NewOrderRepository(db *gorm.DB) *orderRepository {
	return &orderRepository{
		db: db,
	}
}

func (r *orderRepository) Create(order *entity.Order) (*entity.Order, helper.Error) {
	if err := r.db.Create(order).Error; err != nil {
		return nil, helper.NewStatusInternalServerError("failed to create order")
	}
	return order, nil
}

func (r *orderRepository) GetByID(id int) (*entity.Order, helper.Error) {
	var order entity.Order
	if err := r.db.First(&order, id).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, helper.NewStatusNotFoundError("order not found")
		}
		return nil, helper.NewStatusInternalServerError("failed to get order")
	}
	return &order, nil
}

func (r *orderRepository) GetAll() ([]*entity.Order, helper.Error) {
	var orders []*entity.Order
	if err := r.db.Find(&orders).Error; err != nil {
		return nil, helper.NewStatusInternalServerError("failed to get orders")
	}
	return orders, nil
}

func (r *orderRepository) Update(order *entity.Order, id int) (*entity.Order, helper.Error) {
	if err := r.db.Model(&entity.Order{}).Where("id = ?", id).Updates(order).Error; err != nil {
		return nil, helper.NewStatusInternalServerError("failed to update order")
	}
	return order, nil
}

func (r *orderRepository) Delete(id int) (bool, helper.Error) {
	if err := r.db.Delete(&entity.Order{}, id).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return false, helper.NewStatusNotFoundError("order not found")
		}
		return false, helper.NewStatusInternalServerError("failed to delete order")
	}
	return true, nil
}