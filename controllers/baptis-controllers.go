package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	
	"github.com/gorilla/mux"
	"github.com/josepnapitupulu/API_Baptis/models"
	"github.com/josepnapitupulu/API_Baptis/utils"
)

var NewBaptis models.Baptis
// var tmpl * template.Template
// func init(){
// 	tmpl = template.Must(template.ParseFiles("jemaat.html"))
// }

// func Index(w http.ResponseWriter, r *http.Request){
// 	temp, _ := template.ParseFiles("views/jemaat.html")
// 	temp.Execute(w, nil)
// }

// func Create(w http.ResponseWriter, r *http.Request){
// 	temp, _ := template.ParseFiles("views/createJemaat.html")
// 	temp.Execute(w, nil)
// }

func GetBaptis(w http.ResponseWriter, r *http.Request) {
	newBaptiss := models.GetAllBaptiss()
	res, _ := json.Marshal(newBaptiss)
	w.Header().Set("Content-Type", "pkglication/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

func GetBaptisById(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	baptisId := vars["baptisId"]
	NIK, err := strconv.ParseInt(baptisId, 0, 0)
	if err != nil {
		fmt.Println("error while parsing")
	}
	baptisDetails, _ := models.GetBaptisbyId(NIK)
	res, _ := json.Marshal(baptisDetails)
	w.Header().Set("Content-Type", "pkglication/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

func CreateBaptis(w http.ResponseWriter, r *http.Request) {
	CreateBaptis := &models.Baptis{}
	utils.ParseBody(r, CreateBaptis)
	b := CreateBaptis.CreateBaptis()
	res, _ := json.Marshal(b)
	w.Header().Set("Content-Type", "pkglication/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

func DeleteBaptis(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	baptisId := vars["baptisId"]
	ID, err := strconv.ParseInt(baptisId, 0, 0)
	if err != nil {
		fmt.Println("error while parsing")
	}
	baptis := models.DeleteBaptis(ID)
	res, _ := json.Marshal(baptis)
	w.Header().Set("Content-Type", "pkglication/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

func UpdateBaptis(w http.ResponseWriter, r *http.Request) {
    var updateBaptis = &models.Baptis{}
    utils.ParseBody(r, updateBaptis)
    vars := mux.Vars(r)
    baptisId := vars["baptisId"]
    ID, err := strconv.ParseInt(baptisId, 0, 0)
    if err != nil {
        fmt.Println("error while parsing")
    }
    baptisDetails, db := models.GetBaptisbyId(ID)
    if updateBaptis.NIK != "" {
        baptisDetails.NIK = updateBaptis.NIK
    }
    if updateBaptis.NamaAnak != "" {
        baptisDetails.NamaAnak = updateBaptis.NamaAnak
    }
    if updateBaptis.TanggalBaptis != "" {
        baptisDetails.TanggalBaptis = updateBaptis.TanggalBaptis
    }
    if updateBaptis.TanggalLahir != "" {
        baptisDetails.TanggalLahir = updateBaptis.TanggalLahir
    }
    if updateBaptis.NamaAyah != "" {
        baptisDetails.NamaAyah = updateBaptis.NamaAyah
    }
	if updateBaptis.NamaIbu != "" {
        baptisDetails.NamaIbu = updateBaptis.NamaIbu
    }
	if updateBaptis.Alamat != "" {
        baptisDetails.Alamat = updateBaptis.Alamat
    }
    db.Save(&baptisDetails)
    res, _ := json.Marshal(baptisDetails)
    w.Header().Set("Content-Type", "pkglication/json")
    w.WriteHeader(http.StatusOK)
    w.Write(res)
}