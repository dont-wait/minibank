package domain

import "time"

type Customer struct {
	CustomerId string    `json:"customer_id,omitempty" bson:"customer_id"`
	FirstName  string    `json:"first_name,omitempty" bson:"first_name"`
	LastName   string    `json:"last_name,omitempty" bson:"last_name"`
	City       string    `json:"city,omitempty" bson:"city"`
	Phone      string    `json:"phone,omitempty" bson:"phone"`
	Address    string    `json:"address,omitempty" bson:"address"`
	Dob        string    `json:"dob,omitempty" bson:"dob"`
	CreatedAt  time.Time `json:"created_at,omitempty" bson:"created_at"`
	UpdatedAt  time.Time `json:"updated_at,omitempty" bson:"updated_at"`
}
