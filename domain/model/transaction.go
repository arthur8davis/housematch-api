package model

import (
	"github.com/google/uuid"
	"time"
)

type Transaction struct {
	ID         uuid.UUID `json:"id"`
	PropertyID uuid.UUID `json:"property_id"`
	Cost       float64   `json:"cost"`
	Currency   string    `json:"currency"`
	DateVIP    time.Time `json:"date_vip"`
	DatePost   time.Time `json:"date_post"`
	DateUpdate time.Time `json:"date_update"`
	Available  bool      `json:"available"`
	Type       string    `json:"type"`
	DateStart  time.Time `json:"date_start"`
	DateEnd    time.Time `json:"date_end"`
}

type Transactions []Transaction

type TransactionSecondLevel struct {
	ID         uuid.UUID `json:"id"`
	Property   Property  `json:"property"`
	Cost       float64   `json:"cost"`
	Currency   string    `json:"currency"`
	DateVIP    time.Time `json:"date_vip"`
	DatePost   time.Time `json:"date_post"`
	DateUpdate time.Time `json:"date_update"`
	Available  bool      `json:"available"`
	Type       string    `json:"type"`
	DateStart  time.Time `json:"date_start"`
	DateEnd    time.Time `json:"date_end"`
}

type TransactionsSecondLevel []TransactionSecondLevel

type TransactionThirdLevel struct {
	ID         uuid.UUID                      `json:"id"`
	Property   PropertySecondLevelWithoutUser `json:"property"`
	Cost       float64                        `json:"cost"`
	Currency   string                         `json:"currency"`
	DateVIP    time.Time                      `json:"date_vip"`
	DatePost   time.Time                      `json:"date_post"`
	DateUpdate time.Time                      `json:"date_update"`
	Available  bool                           `json:"available"`
	Type       string                         `json:"type"`
	DateStart  time.Time                      `json:"date_start"`
	DateEnd    time.Time                      `json:"date_end"`
}

type TransactionsThirdLevel []TransactionThirdLevel
