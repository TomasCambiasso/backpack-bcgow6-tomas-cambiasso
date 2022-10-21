package domain

type Transaction struct {
	Id               int     `json:"id"`
	Transaction_code string  `json:"transaction_code"`
	Moneda           string  `json:"moneda"`
	Monto            float64 `json:"monto"`
	Emisor           string  `json:"emisor"`
	Receptor         string  `json:"receptor" `
	Transaction_date string  `json:"transaction_date"`
}