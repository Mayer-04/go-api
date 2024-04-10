package models

type Product struct {
	Name        string  `json:"name" bson:"name"`
	Description string  `json:"description" bson:"description"`
	Price       float64 `json:"price" bson:"price"`
	Category    string  `json:"category" bson:"category"`
	Stock       int     `json:"stock" bson:"stock"`
	Image       string  `json:"image" bson:"image"`
}
