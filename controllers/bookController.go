package controllers

import (
	"book-golang/configs"
	"book-golang/helpers"
	"book-golang/models"
	"net/http"

	"github.com/gorilla/mux"
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
	if title == "" {
		helpers.Response(w, 400, "Title is required", nil)
		return
	}

	category := r.Form.Get("category")
	if category == "" {
		helpers.Response(w, 400, "Category is required", nil)
		return
	}


	book := models.Book{
		Title:    title,
		Category: category,
		Stock:    true,
		Borrow:  false,
		Return:  true,
	}

	if err := configs.DB.Create(&book).Error; err != nil {
		helpers.Response(w, 500, err.Error(), nil)
		return
	}

	helpers.Response(w, 200, "Success create book", nil)
}

func GetBook(w http.ResponseWriter, r *http.Request) {
	var book models.Book
	
	params := mux.Vars(r)

	bookID := params["id"]

	if err := configs.DB.First(&book, bookID).Error; err != nil {
		helpers.Response(w, 500, err.Error(), nil)
		return
	}

	helpers.Response(w, 200, "Success get book", book)

}

func DeleteBook(w http.ResponseWriter, r *http.Request) {
	var book models.Book
	
	params := mux.Vars(r)

	bookID := params["id"]

	if err := configs.DB.First(&book, bookID).Error; err != nil {
		helpers.Response(w, 500, err.Error(), nil)
		return
	}

	if err := configs.DB.Delete(&book).Error; err != nil {
		helpers.Response(w, 500, err.Error(), nil)
		return
	}

	helpers.Response(w, 200, "Success delete book", nil)
}

func UpdateBook(w http.ResponseWriter, r *http.Request) {
	var book models.Book
	
	params := mux.Vars(r)

	bookID := params["id"]

	if err := configs.DB.First(&book, bookID).Error; err != nil {
		helpers.Response(w, 500, err.Error(), nil)
		return
	}

	err := r.ParseForm()
	if err != nil {
		helpers.Response(w, 500, err.Error(), nil)
		return
	}

	if title := r.Form.Get("title"); title != "" {
		book.Title = title
	}

	if category := r.Form.Get("category"); category != "" {
		book.Category = category
	}



	if err := configs.DB.Save(&book).Error; err != nil {
		helpers.Response(w, 500, err.Error(), nil)
		return
	}

	helpers.Response(w, 200, "Success update book", nil)
}

func SearchBookByTitle(w http.ResponseWriter, r *http.Request) {
	var books []models.Book

	title := r.URL.Query().Get("title")

	if err := configs.DB.Where("title LIKE ?", "%"+title+"%").Find(&books).Error; err != nil {
		helpers.Response(w, 500, err.Error(), nil)
		return
	}

	helpers.Response(w, 200, "Success get books", books)
}

func FilterBookByCategory(w http.ResponseWriter, r *http.Request) {
	var books []models.Book

	category := r.URL.Query().Get("category")

	if err := configs.DB.Where("category = ?", category).Find(&books).Error; err != nil {
		helpers.Response(w, 500, err.Error(), nil)
		return
	}

	helpers.Response(w, 200, "Success get books", books)
}