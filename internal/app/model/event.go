package model

type Event struct {
	ID            int    `json:"id"`
	Title         string `json:"title"`
	Category      string `json:"category"`
	Description   string `json:"description"`
	StartDatetime string `json:"start_datetime"`
	EndDatetime   string `json:"end-datetime"`
	Price         string `json:"price"`
	Address       string `json:"address"`
	Organizer     string `json:"organizer"`
	Contacts      string `json:"contacts"`
}
