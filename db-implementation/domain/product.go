package domain

type Product struct {
	Id    int     `json:"id"`
	Name  string  `json:"name"`
	Ptype  string  `json:"type"`
	Count int     `json:"count"`
	Price float64 `json:"price"`
}
