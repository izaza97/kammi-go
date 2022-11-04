package diarycontroller

import (
	"encoding/json"
	"kammi/helper"
	"kammi/models"
	"net/http"
	"strconv"
	"time"

	"github.com/gorilla/mux"

)

func GetAllDiary(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	user, err := strconv.ParseInt(vars["user"], 10, 64)
	if err != nil {
		helper.ResponseError(w, http.StatusBadRequest, err.Error())
		return
	}

	var diary []models.Diaryscn
	var diaryu models.Dd
	diaryInput := models.Diary{User: user}
	if err := models.DB.Table("diary-data").Where(&diaryInput, user).Find(&diary).Error; err != nil {
		response := map[string]string{"message": err.Error()}
		helper.ResponseJSON(w, http.StatusInternalServerError, response)
		return
	}

	diaryu.Data = diary
	w.Header().Set("Content-Type", "appication/json")
	json.NewEncoder(w).Encode(diaryu)
}

func CreateDiary(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	errr := r.ParseForm()
	if errr != nil {
		panic(errr)
	}
	//mengambil user dari parameter
	user, err := strconv.ParseInt(vars["user"], 10, 64)
	if err != nil {
		helper.ResponseError(w, http.StatusBadRequest, err.Error())
		return
	}

	// memasukan parameter user ke database dan membuat datetime
	now := time.Now()
	diaryInput := models.Diary{User: user, Time: now.Local()}
	judul := r.Form.Get("judul")
	isi := r.Form.Get("isi")
	diaryInput.Judul = judul
	diaryInput.Isi = isi

	// input ke database
	if err := models.DB.Table("diary-data").Create(&diaryInput).Error; err != nil {
		response := map[string]string{"message": err.Error()}
		helper.ResponseJSON(w, http.StatusInternalServerError, response)
		return
	}

	diaryInput.User = user

	response := map[string]string{"message": "success"}
	helper.ResponseJSON(w, http.StatusOK, response)
}

func GetDiary(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	errr := r.ParseForm()
	if errr != nil {
		panic(errr)
	}
	user, err := strconv.ParseInt(vars["user"], 10, 64)
	if err != nil {
		helper.ResponseError(w, http.StatusBadRequest, err.Error())
		return
	}
	no, err := strconv.ParseInt(vars["no"], 10, 64)
	if err != nil {
		helper.ResponseError(w, http.StatusBadRequest, err.Error())
		return
	}

	var diary []models.Diarys
	var diarys models.Dd2
	diaryInput := models.Diary{User: user, No: no}

	if err := models.DB.Table("diary-data").Where(&diaryInput, user).Find(&diary, no).Error; err != nil {
		response := map[string]string{"message": err.Error()}
		helper.ResponseJSON(w, http.StatusInternalServerError, response)
		return
	}

	diaryInput.User = user
	diarys.Data = diary

	helper.ResponseJSON(w, http.StatusOK, diarys)
}

func UpdateDiary(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	errr := r.ParseForm()
	if errr != nil {
		panic(errr)
	}
	user, err := strconv.ParseInt(vars["user"], 10, 64)
	if err != nil {
		helper.ResponseError(w, http.StatusBadRequest, err.Error())
		return
	}
	//mengambil user dari parameter
	no, err := strconv.ParseInt(vars["no"], 10, 64)
	if err != nil {
		helper.ResponseError(w, http.StatusBadRequest, err.Error())
		return
	}

	// memasukan parameter id ke database
	diaryInput := models.Diary{User: user, No: no}
	judul := r.Form.Get("judul")
	isi := r.Form.Get("isi")
	diaryInput.Judul = judul
	diaryInput.Isi = isi

	// input ke database
	if err := models.DB.Table("diary-data").Where("no = ?", no).Updates(&diaryInput).Error; err != nil {
		response := map[string]string{"message": err.Error()}
		helper.ResponseJSON(w, http.StatusInternalServerError, response)
		return
	}

	diaryInput.No = no

	response := map[string]string{"message": "success"}
	helper.ResponseJSON(w, http.StatusOK, response)
}

func DeleteDiary(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	//mengambil user dari parameter
	no, err := strconv.ParseInt(vars["no"], 10, 64)
	if err != nil {
		helper.ResponseError(w, http.StatusBadRequest, err.Error())
		return
	}

	var diary models.Diary

	// input ke database
	if err := models.DB.Table("diary-data").Where("no = ?", no).Delete(&diary).Error; err != nil {
		response := map[string]string{"message": err.Error()}
		helper.ResponseJSON(w, http.StatusInternalServerError, response)
		return
	}

	response := map[string]string{"message": "success"}
	helper.ResponseJSON(w, http.StatusOK, response)
}
