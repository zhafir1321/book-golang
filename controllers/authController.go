package controllers

import (
	"book-golang/configs"
	"book-golang/helpers"
	"book-golang/models"
	"net/http"
)

func Register(w http.ResponseWriter, r *http.Request) {

	err := r.ParseForm()
	if err != nil {
		helpers.Response(w, 500, err.Error(), nil)
		return
	}

	name := r.Form.Get("name")
	email := r.Form.Get("email")
	password := r.Form.Get("password")
	role := "user"

	passwordHash, err := helpers.HashPassword(password)
	if err != nil {
		helpers.Response(w, 500, err.Error(), nil)
		return
	}

	user := models.User{
		Name:    name,
		Email:   email,
		Password: passwordHash,
		Role:    role,
	}

	if err := configs.DB.Create(&user).Error; err != nil {
		helpers.Response(w, 500, err.Error(), nil)
		return
	}

	helpers.Response(w, 200, "Success create user", nil)
}

func Login(w http.ResponseWriter, r *http.Request) {
	err := r.ParseForm()
	if err != nil {
		helpers.Response(w, 500, err.Error(), nil)
		return
	}

	email := r.Form.Get("email")
	password := r.Form.Get("password")

	var user models.User
	if err := configs.DB.First(&user, "email = ?", email).Error; err != nil {
		helpers.Response(w, 404, "Invalid Email or Password", nil)
		return
	}

	if err := helpers.VerifyPassword(user.Password, password); err != nil {
		helpers.Response(w, 404, "Invalid Email or Password", nil)
		return
	}

	token, err := helpers.CreateToken(&user)

	if err != nil {
		helpers.Response(w, 500, err.Error(), nil)
		return
	}

	helpers.Response(w, 200, "Success login", token)

	
}