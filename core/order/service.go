package order

import (
	"database/sql"
	"log"
)

type UseCase interface {
	GetAll() ([]*Order, error)
	Get(ID int64) (Order, error)
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


func (s *Service) Get(ID int64) (Order, error) {
	var order Order
	sqlStatement := `SELECT * FROM order where id = $1`
	err := s.DB.QueryRow(sqlStatement, ID).Scan(&order.ID, &order.NumberOrder, &order.ClientName,
		&order.Address, &order.EstimatedTimeOfArrival, &order.Status, &order.Pizzas)
	if err != nil {
		log.Fatal(err)
	}
	return order, nil
}

func (s *Service) Store(o *Order) error {
	tx, err := s.DB.Begin()
	if err != nil {
		return err
	}
	stmt, err := tx.Prepare("INSERT INTO public.order (number_order, client_name, address, pizzas, estimated_time_of_arrival, status) VALUES ($1, $2, $3, $4, $5, $6)")
	if err != nil {
		return err
	}
	defer stmt.Close()
	_, err = stmt.Exec(o.NumberOrder,o.ClientName, o.Address, o.Pizzas, o.EstimatedTimeOfArrival, o.Status)
	if err != nil {
		tx.Rollback()
		return err
	}
	tx.Commit()
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