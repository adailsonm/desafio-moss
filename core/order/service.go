package order

import (
	"gorm.io/gorm"
	"time"
)

type UseCase interface {
	GetAll() ([]*Order, error)
	Get(ID int64) (*Order, error)
	Store(o *Order) error
	Update(o *Order) error
	Remove(ID int64) error
}

type Service struct {
	DB *gorm.DB
}

func (s *Service) GetAll() ([]*Order, error) {
	panic("implement me")
}


func (s *Service) Get(ID int64) (*Order, error) {
	panic("implement me")
}

func (s *Service) Store(o *Order) error {
	order := Order{
		NumberOrder: o.NumberOrder,
		ClientName: o.ClientName,
		Pizzas: o.Pizzas,
		Price: o.Price,
		EstimatedTimeOfArrival: o.EstimatedTimeOfArrival,
		LastUpdate: time.Now(),
		Status: o.Status,
	}

	result := s.DB.Create(&order)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

func (s *Service) Update(o *Order) error {
	panic("implement me")
}

func (s *Service) Remove(ID int64) error {
	panic("implement me")
}

func NewService(db *gorm.DB) *Service {
	return &Service{
		DB: db,
	}
}