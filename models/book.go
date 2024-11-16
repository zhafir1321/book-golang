package models

type Book struct {
	ID       uint `gorm:"primaryKey"`
	Title    string
	Category string
	Stock    bool
	Borrow   bool
	Return   bool
}

type CreateBook struct {
	Title    string `json:"title"`
	Category string `json:"category"`
	Stock    int    `json:"stock"`
}