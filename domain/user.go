package domain

type User struct {
	Name     string  `json:"name"`
	Username string  `json:"username"`
	Address  Address `json:"address"`
}

type Address struct {
	No      int    `json:"no"`
	Address string `json:"address"`
}
