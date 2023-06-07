package model

type Ticket struct {
	ID           string `json:"id"`
	Price        string `json:"price"`
	PurchaseDate string `json:"purchase_date"`
	User         string `json:"user_id"`
	Event        string `json:"event_id"`
}
