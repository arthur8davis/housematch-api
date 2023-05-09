package model

import (
	"github.com/google/uuid"
	"time"
)

type ExchangeRate struct {
	ID            uuid.UUID `json:"id"`
	Date          time.Time `json:"date"`
	Currency      string    `json:"currency"`
	PurchasePrice float64   `json:"purchase_price"`
	SellingPrice  float64   `json:"selling_price"`
}
