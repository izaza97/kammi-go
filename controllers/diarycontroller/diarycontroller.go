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

	diary := []models.Diaryscn{}
	diaryInput := models.Diary{User: user}
	if err := models.DB.Table("diary-data").Where(&diaryInput, user).Find(&diary).Error; err != nil {
		response := map[string]string{"message": err.Error()}
		helper.ResponseJSON(w, http.StatusInternalServerError, response)
		return
	}

	diaryInput.User = user

	helper.ResponseJSON(w, http.StatusOK, diary)
}

func CreateDiary(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	//mengambil user dari parameter
	user, err := strconv.ParseInt(vars["user"], 10, 64)
	if err != nil {
		helper.ResponseError(w, http.StatusBadRequest, err.Error())
		return
	}

	// memasukan parameter user ke database dan membuat datetime
	diaryInput := models.Diary{User: user, Time: time.Now()}
	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&diaryInput); err != nil {
		helper.ResponseError(w, http.StatusInternalServerError, err.Error())
		return
	}

	defer r.Body.Close()

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

	var diary models.Diaryscn
	diaryInput := models.Diary{User: user}

	if err := models.DB.Table("diary-data").Where(&diaryInput, user).Find(&diary, no).Error; err != nil {
		response := map[string]string{"message": err.Error()}
		helper.ResponseJSON(w, http.StatusInternalServerError, response)
		return
	}

	diaryInput.User = user
	diary.No = no

	helper.ResponseJSON(w, http.StatusOK, diary)
}

func UpdateDiary(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	//mengambil user dari parameter
	no, err := strconv.ParseInt(vars["no"], 10, 64)
	if err != nil {
		helper.ResponseError(w, http.StatusBadRequest, err.Error())
		return
	}

	// memasukan parameter id ke database
	diaryInput := models.Diary{No: no}
	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&diaryInput); err != nil {
		helper.ResponseError(w, http.StatusInternalServerError, err.Error())
		return
	}

	defer r.Body.Close()

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
