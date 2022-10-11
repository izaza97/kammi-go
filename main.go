package main

import (
	"kammi/controllers/ahd"
	"kammi/controllers/authcontroller"
	"kammi/controllers/diarycontroller"
	"kammi/controllers/productcontroller"
	"kammi/middlewares"
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
	r.HandleFunc("/logout", authcontroller.Logout).Methods("GET")
	r.HandleFunc("/ashusna", ahd.Ashusna).Methods("GET")

	api := r.PathPrefix("/api").Subrouter()
	api.HandleFunc("/products", productcontroller.Index).Methods("GET")
	api.HandleFunc("/sign/{id}", authcontroller.Show).Methods("GET")
	api.HandleFunc("/sign/{id}/update", authcontroller.Updateuser).Methods("POST")
	api.HandleFunc("/sign/{user}/diary", diarycontroller.GetAllDiary).Methods("GET")
	api.HandleFunc("/sign/{user}/diary/create", diarycontroller.CreateDiary).Methods("POST")
	api.HandleFunc("/sign/{user}/diary/{no}", diarycontroller.GetDiary).Methods("GET")
	api.HandleFunc("/sign/{user}/diary/{no}/update", diarycontroller.UpdateDiary).Methods("POST")
	api.HandleFunc("/sign/{user}/diary/{no}/delete", diarycontroller.DeleteDiary).Methods("POST")
	api.Use(middlewares.JWTMiddleware)

	log.Fatal(http.ListenAndServe(":3000", r))
}
