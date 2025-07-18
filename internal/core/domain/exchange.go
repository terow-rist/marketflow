package domain

import "time"

type PriceUpdate struct {
	Symbol    string    `json:"symbol"`
	Price     float64   `json:"price"`
	Timestamp time.Time `json:"timestamp"`
	Exchange  string
}
