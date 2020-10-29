package models

// Person ...
type Person struct {
	UserID    string    `json:"user_id"`
	FirstName string    `json:"first_name"`
	LastName  string    `json:"last_name"`
	Addresses []Address `json:"addresses"`
}

// InsertInfo ...
type InsertInfo struct {
	FirstName string    `json:"first_name"`
	LastName  string    `json:"last_name"`
	Addresses []Address `json:"addresses"`
}

// Address ...
type Address struct {
	Address string `json:"address"`
}