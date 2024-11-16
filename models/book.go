package models

type Book struct {
	ID       uint `gorm:"primaryKey"`
	Title    string
	Category string
	Stock    int
}

type CreateBook struct {
	Title    string `json:"title"`
	Category string `json:"category"`
	Stock    int    `json:"stock"`
}