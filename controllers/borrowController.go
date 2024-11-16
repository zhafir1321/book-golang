package controllers

import (
	"book-golang/configs"
	"book-golang/helpers"
	"book-golang/models"
	"net/http"
	"strconv"
	"time"

	"github.com/gorilla/mux"
	"gorm.io/gorm"
)


func CreateBorrow(w http.ResponseWriter, r *http.Request) {

	bookParams := mux.Vars(r)
	bookID, err := strconv.ParseUint(bookParams["bookId"], 10, 64)
	if err != nil {
		helpers.Response(w, 400, "Invalid book ID", nil)
		return
	}

	userInfo, ok := r.Context().Value("userInfo").(*helpers.MyCustomClaims)
	if !ok {
		helpers.Response(w, 400, "Unauthorized", nil)
		return
	}

	borrow := models.BorrowReturn{
		UserID: userInfo.ID,
		BookID: uint(bookID),
		BorrowDate: time.Now(),
		ReturnDate: time.Now().AddDate(0, 0, 7),
		Status: "borrowing",
	}

	if err := configs.DB.Create(&borrow).Error; err != nil {
		helpers.Response(w, 500, err.Error(), nil)
		return
	}

	book := models.Book{}
	if err := configs.DB.First(&book, "id = ?", bookID).Error; err != nil {
		helpers.Response(w, 500, err.Error(), nil)
		return
	}

	book.Stock = false
	book.Borrow = true
	book.Return = false

	if err := configs.DB.Save(&book).Error; err != nil {
		helpers.Response(w, 500, err.Error(), nil)
		return
	}



	helpers.Response(w, 200, "Success borrow book", nil)

}


func ReturnBook(w http.ResponseWriter, r *http.Request) {
	
	bookParams := mux.Vars(r)
	bookID, err := strconv.ParseUint(bookParams["bookId"], 10, 64)
	if err != nil {
		helpers.Response(w, 400, "Invalid book ID", nil)
		return
	}

	userInfo, ok := r.Context().Value("userInfo").(*helpers.MyCustomClaims)
	if !ok {
		helpers.Response(w, 400, "Unauthorized", nil)
		return
	}

	borrow := models.BorrowReturn{}
	if err := configs.DB.First(&borrow, "user_id = ? AND book_id = ?", userInfo.ID, bookID).Error; err != nil {
		helpers.Response(w, 500, "Book not found", nil)
		return
	}


	borrow.Status = "returned"

	if err := configs.DB.Save(&borrow).Error; err != nil {
		helpers.Response(w, 500, err.Error(), nil)
		return
	}

	book := models.Book{}
	if err := configs.DB.First(&book, "id = ?", bookID).Error; err != nil {
		helpers.Response(w, 500, err.Error(), nil)
		return
	}

	book.Stock = true
	book.Borrow = false
	book.Return = true

	if err := configs.DB.Save(&book).Error; err != nil {
		helpers.Response(w, 500, err.Error(), nil)
		return
	}

	helpers.Response(w, 200, "Success return book", nil)

}

func GetBorrows(w http.ResponseWriter, r *http.Request) {

	userInfo, ok := r.Context().Value("userInfo").(*helpers.MyCustomClaims)
	if !ok {
		helpers.Response(w, 400, "Unauthorized", nil)
		return
	}

	borrows := []models.BorrowReturn{}
	if err := configs.DB.Preload("User", func(db *gorm.DB) *gorm.DB {
        return db.Omit("password")
    }).Preload("Book").Find(&borrows, "user_id = ?", userInfo.ID).Error; err != nil {
		helpers.Response(w, 500, err.Error(), nil)
		return
	}

	helpers.Response(w, 200, "Success get borrows", borrows)

}