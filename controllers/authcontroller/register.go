package authcontroller

import (
	"crypto/sha256"
	"encoding/hex"
	"kammi/helper"
	"kammi/models"
	"net/http"

)

//register controller
func Register(w http.ResponseWriter, r *http.Request) {

	// mengambil inputan json
	var userInput models.User
	err := r.ParseForm()
	if err != nil {
		panic(err)
	}
	name := r.Form.Get("name")
	username := r.Form.Get("username")
	email := r.Form.Get("email")
	password := r.Form.Get("password")
	passwordconfirm := r.Form.Get("passwordconfirm")
	img := r.Form.Get("img")
	userInput.Name = name
	userInput.Username = username
	userInput.Email = email
	userInput.Password = password
	userInput.Img = img

	// amankan pass menggunakan sha256
	pass := sha256.New()
	pass.Write([]byte(userInput.Password))
	shapass := pass.Sum(nil)
	userInput.Password = hex.EncodeToString(shapass)

	// ambil data user berdasarkan username
	var user models.User
	if err := models.DB.Table("user-data").Where("username = ?", userInput.Username).First(&user).Error; err != nil {
		switch err {
		default:
			// insert ke database
			if password == passwordconfirm {
				models.DB.Table("user-data").Create(&userInput)
				response := map[string]string{"message": "Success"}
				helper.ResponseJSON(w, http.StatusInternalServerError, response)
				return
			} else {
				response := map[string]string{"message": "Failed"}
				helper.ResponseJSON(w, http.StatusInternalServerError, response)
				return
			}
		}
	} else {
		response := map[string]string{"message": "Failed"}
		helper.ResponseJSON(w, http.StatusUnauthorized, response)
		return
	}
}
