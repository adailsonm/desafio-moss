package order

import (
	"database/sql"
	"log"
)

type UseCase interface {
	GetAll() ([]*Order, error)
	Get(ID int64) (*Order, error)
	Store(o *Order) error
	Update(o *Order) error
	Remove(ID int64) error
}

type Service struct {
	DB *sql.DB
}

func (s *Service) GetAll() ([]*Order, error) {
	panic("implement me")
}


func (s *Service) Get(ID int64) (*Order, error) {
	panic("implement me")
}

func (s *Service) Store(o *Order) error {
	sqlStatement := `INSERT INTO order (number_order, client_name, address, pizzas, estimated_time_of_arrival, status)
		VALUES ($1, $2, $3, $4, $5, $6)`
	_, err := s.DB.Exec(sqlStatement, o.NumberOrder,o.ClientName, o.Address, o.Pizzas, o.EstimatedTimeOfArrival, o.Status)
	if err != nil {
		log.Fatal(err)
	}
	return nil
}

func (s *Service) Update(o *Order) error {
	panic("implement me")
}

func (s *Service) Remove(ID int64) error {
	panic("implement me")
}

func NewService(db *sql.DB) *Service {
	return &Service{
		DB: db,
	}
}