package model

type CustomerShipping struct {
	Id        string           `json:"id"`
	Name      string           `json:"name"`
	Address   *CustomerAddress `json:"address"`
	CreatedAt int64            `json:"created_at"`
}

type CustomerAddress struct {
	Id        string  `json:"id"`
	City      string  `json:"city"`
	Street    string  `json:"street"`
	Longitude float64 `json:"longitude"`
	Latitude  float64 `json:"latitude"`
}
