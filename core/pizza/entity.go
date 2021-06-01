package pizza

type Pizza struct {
	ID    int64     `json:"id"`
	Name  string  `json:"name"`
	Ingredients  []string `json:"ingredients"`
	Price      float64       `json:"price"`
}

