package main

import (
	"kammi/controllers/ahd"
	"kammi/controllers/authcontroller"
	"kammi/controllers/dhariancontroller"
	"kammi/controllers/diarycontroller"
	"kammi/controllers/mssgcontroller"
	"kammi/controllers/productcontroller"
	"kammi/models"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {

	models.ConnectDatabase()
	r := mux.NewRouter()

	r.HandleFunc("/login", authcontroller.Login).Methods("POST")
	r.HandleFunc("/register", authcontroller.Register).Methods("POST")
	r.HandleFunc("/Img", authcontroller.Imgp).Methods("GET")

	r.HandleFunc("/ashusna", ahd.Ashusna).Methods("GET")
	r.HandleFunc("/dharian", dhariancontroller.KDharian).Methods("GET")
	r.HandleFunc("/dharian/{katagori}", dhariancontroller.Show).Methods("GET")

	r.HandleFunc("/products", productcontroller.Index).Methods("GET")

	r.HandleFunc("/sign/{id}", authcontroller.Show).Methods("GET")
	r.HandleFunc("/sign/{id}/update", authcontroller.Updateprofile).Methods("POST")
	r.HandleFunc("/sign/{id}/updateuser", authcontroller.Updateusername).Methods("POST")
	r.HandleFunc("/sign/{id}/sun", authcontroller.Showun).Methods("GET")
	r.HandleFunc("/sign/{id}/updateun", authcontroller.Updateusername).Methods("POST")
	r.HandleFunc("/sign/{id}/updatepw", authcontroller.Updatepw).Methods("POST")

	r.HandleFunc("/sign/{user}/diary", diarycontroller.GetAllDiary).Methods("GET")
	r.HandleFunc("/sign/{user}/diary/create", diarycontroller.CreateDiary).Methods("POST")
	r.HandleFunc("/sign/{user}/diary/{no}", diarycontroller.GetDiary).Methods("GET")
	r.HandleFunc("/sign/{user}/diary/{no}/update", diarycontroller.UpdateDiary).Methods("POST")
	r.HandleFunc("/sign/{user}/diary/{no}/delete", diarycontroller.DeleteDiary).Methods("POST")

	r.HandleFunc("/sign/{id}/message", mssgcontroller.Message).Methods("POST")

	log.Fatal(http.ListenAndServe(":3000", r))
}
