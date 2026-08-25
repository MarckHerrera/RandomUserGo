package models

/* 1 */
type RandomUser struct {
	Results []User `json:"results"`
}

/* 2 del RandomUser */
type User struct {
	Gender string `json:"gender"`
	Name Name `json:"name"`
	Location Location `json:"location"`
	Email string `json:"email"`
	Phone string `json:"phone"`
}

/* 3 del User */
type Name struct {
	Title string `json:"title"`
	First string `json:"first"`
	Last string `json:"last"`
}

/* 3 del User */
type Location struct {
	Street Street `json:"street"`
	City string `json:"city"`
	State string `json:"state"`
	Country string `json:"country"`
}


/* 4 del Location */
type Street struct {
	Number int `json:"number"`
	Name string `json:"name"`
}