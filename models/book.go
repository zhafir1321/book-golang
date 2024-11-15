package models

type Book struct {
	ID        uint `gorm:"primaryKey"`
	Title     string
	Category  string
	Stock     int
	Borrowing bool `gorm:"default:false"`
	Returning bool `gorm:"default:true"`
}