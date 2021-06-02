package order

import (
	"encoding/json"
	"time"
)

type Order struct {
	ID    uint     `json:"id"`
	NumberOrder    string     `json:"number_order"`
	ClientName    string     `json:"client_name"`
	Address    string `json:"address"`
	Pizzas    json.RawMessage     `json:"pizzas"`
	EstimatedTimeOfArrival  time.Time  `json:"estimated_time_of_arrival"`
	LastUpdate  time.Time `json:"last_update"`
	Status OrderStatus `json:"status"`
}

type OrderStatus int

const (
	Pending = iota + 1
	Preparing
	OnRoute
	Concluded
	Canceled
)

func (o OrderStatus) String() string {
	switch o {
	case Pending:
		return "Pendente"
	case Preparing:
		return "Preparando"
	case OnRoute:
		return "Em rota"
	case Concluded:
		return "Concluido"
	case Canceled:
		return "Cancelado"
	}
	return "Unknown"
}

