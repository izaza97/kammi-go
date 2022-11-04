package authcontroller

import (
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"
	"kammi/helper"
	"kammi/models"
	"log"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"gorm.io/gorm"
)

//show userbyid controller
func Show(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)
	id, err := strconv.ParseInt(vars["id"], 10, 64)
	if err != nil {
		helper.ResponseError(w, http.StatusBadRequest, err.Error())
		return
	}

	var user []models.UserS
	var useri models.Ud
	if err := models.DB.Table("user-data").Select("`user-data`.id, `user-data`.`name`, `user-data`.`username`, `user-data`.`email`, `user-data`.`password`,`user-img`.`img`").Joins("INNER JOIN `user-img` ON `user-data`.`img` = `user-img`.`id`").First(&user, id).Error; err != nil {
		switch err {
		case gorm.ErrRecordNotFound:
			helper.ResponseError(w, http.StatusNotFound, "user tidak ditemukan")
			return
		default:
			helper.ResponseError(w, http.StatusInternalServerError, err.Error())
			return
		}
	}
	useri.Data = user
	w.Header().Set("Content-Type", "appication/json")
	json.NewEncoder(w).Encode(useri)

}

func Showun(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)
	id, err := strconv.ParseInt(vars["id"], 10, 64)
	if err != nil {
		helper.ResponseError(w, http.StatusBadRequest, err.Error())
		return
	}

	var user []models.UU
	var useri models.Un
	if err := models.DB.Table("user-data").First(&user, id).Error; err != nil {
		switch err {
		case gorm.ErrRecordNotFound:
			helper.ResponseError(w, http.StatusNotFound, "user tidak ditemukan")
			return
		default:
			helper.ResponseError(w, http.StatusInternalServerError, err.Error())
			return
		}
	}
	useri.Data = user
	w.Header().Set("Content-Type", "appication/json")
	json.NewEncoder(w).Encode(useri)
}

//update user controller
func Updateprofile(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)
	id, err := strconv.ParseInt(vars["id"], 10, 64)
	if err != nil {
		helper.ResponseError(w, http.StatusBadRequest, err.Error())
		return
	}

	var userInput models.UserU
	errr := r.ParseForm()
	if errr != nil {
		panic(err)
	}
	name := r.Form.Get("name")
	username := r.Form.Get("username")
	email := r.Form.Get("email")
	img := r.Form.Get("img")
	userInput.Username = username
	userInput.Name = name
	userInput.Email = email
	userInput.Img = img

	defer r.Body.Close()

	//
	var user models.UU
	if err := models.DB.Table("user-data").Where("username = ?", userInput.Username).First(&user).Error; err != nil {
		switch err {
		default:
			// insert ke database
			models.DB.Table("user-data").Where("id = ?", id).Updates(&userInput)
			response := map[string]string{"message": "success"}
			helper.ResponseJSON(w, http.StatusInternalServerError, response)
			return
		}
	} else {
		response := map[string]string{"message": "Username already taken"}
		helper.ResponseJSON(w, http.StatusUnauthorized, response)
		return
	}
	// if models.DB.Table("user-data").Where("id = ?", id).Updates(&userInput).RowsAffected == 0 {
	// 	helper.ResponseError(w, http.StatusBadRequest, "Tidak dapat mengupdate user")
	// 	return
	// }

	// userInput.Id = id

	// response := map[string]string{"message": "success"}
	// helper.ResponseJSON(w, http.StatusOK, response)
}

func Updateusername(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)
	id, err := strconv.ParseInt(vars["id"], 10, 64)
	if err != nil {
		helper.ResponseError(w, http.StatusBadRequest, err.Error())
		return
	}

	var userInput models.UU
	errr := r.ParseForm()
	if errr != nil {
		panic(err)
	}
	username := r.Form.Get("username")
	userInput.Username = username

	defer r.Body.Close()

	//
	var user models.UU
	if err := models.DB.Table("user-data").Where("username = ?", userInput.Username).First(&user).Error; err != nil {
		switch err {
		default:
			// insert ke database
			models.DB.Table("user-data").Where("id = ?", id).Updates(&userInput)
			response := map[string]string{"message": "success"}
			helper.ResponseJSON(w, http.StatusInternalServerError, response)
			return
		}
	} else {
		response := map[string]string{"message": "Username already taken"}
		helper.ResponseJSON(w, http.StatusUnauthorized, response)
		return
	}
}

func Updatepw(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)
	id, err := strconv.ParseInt(vars["id"], 10, 64)
	if err != nil {
		helper.ResponseError(w, http.StatusBadRequest, err.Error())
		return
	}

	var userInput models.UP
	errr := r.ParseForm()
	if errr != nil {
		panic(err)
	}
	password := r.Form.Get("password")
	passwordconfirm := r.Form.Get("passwordconfirm")
	userInput.Password = password

	defer r.Body.Close()

	//
	pass := sha256.New()
	pass.Write([]byte(userInput.Password))
	shapass := pass.Sum(nil)
	userInput.Password = hex.EncodeToString(shapass)

	if password == passwordconfirm {
		models.DB.Table("user-data").Where("id = ?", id).Updates(&userInput)
		response := map[string]string{"message": "success"}
		helper.ResponseJSON(w, http.StatusInternalServerError, response)
		return
	} else {
		response := map[string]string{"message": "password not same"}
		helper.ResponseJSON(w, http.StatusInternalServerError, response)
		return
	}
}

func Imgp(w http.ResponseWriter, r *http.Request) {
	img := []models.Imgp{}
	var response models.Ip
	result := models.DB.Table("user-img").Scan(&img).Error
	if result != nil {
		log.Print(result.Error())
	}
	response.Data = img
	w.Header().Set("Content-Type", "appication/json")
	json.NewEncoder(w).Encode(response)
}
