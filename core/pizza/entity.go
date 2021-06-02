package pizza

import "encoding/json"

type Pizza struct {
	ID    int64     `json:"id"`
	Name  string  `json:"name"`
	Ingredients json.RawMessage `json:"ingredients,omitempty"`
	Price      float64       `json:"price"`
}
