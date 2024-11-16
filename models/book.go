package models

type Book struct {
	ID       uint   `gorm:"primaryKey"`
	Title    string `gorm:"not null"`
	Category string `gorm:"not null"`
	Stock    bool   `gorm:"not null"`
	Borrow   bool   `gorm:"not null"`
	Return   bool   `gorm:"not null"`
}

type CreateBook struct {
	Title    string `json:"title"`
	Category string `json:"category"`
	Stock    int    `json:"stock"`
}