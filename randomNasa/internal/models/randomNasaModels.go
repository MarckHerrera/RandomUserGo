package models

type RandomUser struct {
	Results []User `json:"results"`
}

type User struct {
	Gender string `json:"gender"`
	Name Name `json:"name"`
	Location Location `json:"location"`
	Email string `json:"email"`
	Phone string `json:"phone"`
}

type Name struct {
	First string `json:"first"`
	Last string `json:"last"`
}

type Location struct {
	Street Street `json:"street"`
	City string `json:"city"`
	State string `json:"state"`
	Country string `json:"country"`
}

type Street struct {
	Number int `json:"number"`
	Name string `json:"name"`
}