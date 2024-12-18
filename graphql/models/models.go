package models

import "time"

type Account struct {
	ID string `json:"id"`
	Name string `json:"name"`
	Password string `json:"password"`
	Email string `json:"email"`
	Orders []Order `json:"orders"`
	ShopNames []ShopName `json:"shopnames"`
}

type ShopName struct {
	Shopname string `json:"shopname"`
}

type Order struct {
	ID string `json:"id"`
	Amount float64 `json:"amount"`
	AccountID string `json:"accountId"`
	CreatedAt time.Time `json:"createdAt"`
	Description string `json:"description"`
	LineItems []OrderLineItem `json:"lineItems"`
}

type OrderLineItem struct {
	ID string `json:"id"`
	Amount float64 `json:"amount"`
	Description string `json:"description"`
}

type AccountInput struct {
	Name string `json:"name"`
}
type OrderInput struct {
	AccountID string `json:"accountId"`
	Amount float64 `json:"amount"`
	Description string `json:"description"`
}