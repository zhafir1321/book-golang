package controllers

import (
	"book-golang/configs"
	"book-golang/helpers"
	"book-golang/models"
	"net/http"
	"strconv"
)



func GetBooks(w http.ResponseWriter, r *http.Request) {
	var books []models.Book

	if err := configs.DB.Find(&books).Error; err != nil {
		helpers.Response(w, 500, err.Error(), nil)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	helpers.Response(w, 200, "Success get books", books)
}

func CreateBook(w http.ResponseWriter, r *http.Request) {
	err := r.ParseForm()
	if err != nil {
		helpers.Response(w, 500, err.Error(), nil)
		return
	}

	title := r.Form.Get("title")
	category := r.Form.Get("category")
	stock := r.Form.Get("stock")
	stockInt, err := strconv.Atoi(stock)
	if err != nil {
		helpers.Response(w, 500, "Invalid stock value", nil)
		return
	}

	book := models.Book{
		Title:    title,
		Category: category,
		Stock:    stockInt,
	}

	if err := configs.DB.Create(&book).Error; err != nil {
		helpers.Response(w, 500, err.Error(), nil)
		return
	}

	helpers.Response(w, 200, "Success create book", nil)
	


}