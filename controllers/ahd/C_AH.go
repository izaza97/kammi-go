package ahd

import (
	"encoding/json"
	"kammi/models"
	"log"
	"net/http"

)

func Ashusna(w http.ResponseWriter, r *http.Request) {
	husna := []models.AH{}
	var response models.R_ah
	result := models.DB.Table("asmaul-husna-data").Scan(&husna).Error
	if result != nil {
		log.Print(result.Error())
	}
	response.Data = husna
	w.Header().Set("Content-Type", "appication/json")
	json.NewEncoder(w).Encode(response)
}
