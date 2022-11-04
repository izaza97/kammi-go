package dhariancontroller

import (
	"encoding/json"
	"kammi/helper"
	"kammi/models"
	"log"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"

)

func KDharian(w http.ResponseWriter, r *http.Request) {
	kdh := []models.Kdh{}
	var response models.Ddh2
	result := models.DB.Table("doa-harian-kategori").Scan(&kdh).Error
	if result != nil {
		log.Print(result.Error())
	}
	response.Data = kdh
	w.Header().Set("Content-Type", "appication/json")
	json.NewEncoder(w).Encode(response)
}

//show doa harianbyid controller
func Show(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	katagori, err := strconv.ParseInt(vars["katagori"], 10, 64)
	if err != nil {
		helper.ResponseError(w, http.StatusBadRequest, err.Error())
		return
	}

	var dh []models.Dh
	var Ddh models.Ddh
	dhInput := models.Dh{Katagori: int(katagori)}
	if err := models.DB.Table("doa-harian-data").Where(&dhInput, katagori).Find(&dh).Error; err != nil {
		response := map[string]string{"message": err.Error()}
		helper.ResponseJSON(w, http.StatusInternalServerError, response)
		return
	}

	Ddh.Data = dh
	w.Header().Set("Content-Type", "appication/json")
	json.NewEncoder(w).Encode(Ddh)
}
