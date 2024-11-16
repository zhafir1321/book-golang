package models

import "time"

type BorrowReturn struct {
	ID         uint `gorm:"primaryKey"`
	BorrowDate time.Time `gorm:"not null"`
	ReturnDate time.Time `gorm:"not null"`
	UserID     uint
	BookID     uint
	Status	 string

	User User `gorm:"foreignKey:UserID;constraint:OnDelete:CASCADE"`
	Book Book `gorm:"foreignKey:BookID;constraint:OnDelete:CASCADE"`
}