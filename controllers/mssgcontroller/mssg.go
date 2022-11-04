package mssgcontroller

import (
	"kammi/helper"
	"kammi/models"
	"log"
	"net/http"
	"strconv"
	"time"

	"github.com/gorilla/mux"
	"github.com/jinzhu/gorm"
	"gopkg.in/gomail.v2"

)

const CONFIG_SMTP_HOST = "smtp.gmail.com"
const CONFIG_SMTP_PORT = 587
const CONFIG_SENDER_NAME = ""
const CONFIG_AUTH_EMAIL = "baritasetiawati@gmail.com"
const CONFIG_AUTH_PASSWORD = "zvbminayuiimuybq"

func Message(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	errr := r.ParseForm()
	if errr != nil {
		panic(errr)
	}
	//mengambil user dari parameter
	id, err := strconv.ParseInt(vars["id"], 10, 64)
	if err != nil {
		helper.ResponseError(w, http.StatusBadRequest, err.Error())
		return
	}

	// memasukan parameter user ke database dan membuat datetime
	now := time.Now()
	mssgInput := models.Mail{Id: id, Datetime: now.Local()}
	subject := r.Form.Get("subject")
	message := r.Form.Get("message")
	mssgInput.Message = message

	var user models.UserS
	// var useri models.Un
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

	// input ke database
	if err := models.DB.Table("user-message").Create(&mssgInput).Error; err != nil {
		response := map[string]string{"message": err.Error()}
		helper.ResponseJSON(w, http.StatusInternalServerError, response)
		return
	}
	mssgInput.Id = id

	response := map[string]string{"message": "success"}
	helper.ResponseJSON(w, http.StatusOK, response)
	mailer := gomail.NewMessage()
	mailer.SetHeader("From", user.Username)
	mailer.SetHeader("To", "kammi.fmn@gmail.com")
	mailer.SetHeader("Subject", subject)
	mailer.SetBody("text/html", message)

	dialer := gomail.NewDialer(
		CONFIG_SMTP_HOST,
		CONFIG_SMTP_PORT,
		CONFIG_AUTH_EMAIL,
		CONFIG_AUTH_PASSWORD,
	)

	err2 := dialer.DialAndSend(mailer)
	if err2 != nil {
		log.Fatal(err.Error())
	}

	log.Println("Mail sent!")
}
