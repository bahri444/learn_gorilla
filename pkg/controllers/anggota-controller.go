package controllers

import (
	"encoding/json"
	"fmt"
	"learn_gorilla/pkg/models"
	"learn_gorilla/pkg/utils"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

var NewAnggota models.Anggota

func GetAnggota(w http.ResponseWriter, r *http.Request) {
	newAnggotas := models.GetAllAnggota()
	res, _ := json.Marshal(newAnggotas)
	w.Header().Set("Content-Type", "pkglication/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

func GetAnggotaById(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	anggotaId := vars["anggotaId"]
	ID, err := strconv.ParseInt(anggotaId, 0, 0)
	if err != nil {
		fmt.Println("error while parsing")
	}
	anggotaDetails, _ := models.GetAnggotaById(ID)
	res, _ := json.Marshal(anggotaDetails)
	w.Header().Set("Content-Type", "pkglication/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

func CreateAnggota(w http.ResponseWriter, r *http.Request) {
	CreateAnggota := &models.Anggota{}
	utils.ParseBody(r, CreateAnggota)
	a := CreateAnggota.CreateAnggota()
	res, _ := json.Marshal(a)
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

func DeleteAnggota(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	anggotaId := vars["anggotaId"]
	ID, err := strconv.ParseInt(anggotaId, 0, 0)
	if err != nil {
		fmt.Println("error while parsing")
	}
	anggota := models.DeleteAnggota(ID)
	res, _ := json.Marshal(anggota)
	w.Header().Set("Content-Type", "pkglication/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

func UpdateAnggota(w http.ResponseWriter, r *http.Request) {
	var updateAnggota = &models.Anggota{}
	utils.ParseBody(r, updateAnggota)
	vars := mux.Vars(r)
	anggotaId := vars["anggotaId"]
	ID, err := strconv.ParseInt(anggotaId, 0, 0)
	if err != nil {
		fmt.Println("error while parsing")
	}
	anggotaDetails, db := models.GetAnggotaById(ID)
	if updateAnggota.Nim != "" {
		anggotaDetails.Nim = updateAnggota.Nim
	}

	if updateAnggota.Nama != "" {
		anggotaDetails.Nama = updateAnggota.Nama
	}

	if updateAnggota.JenisKelamin != "" {
		anggotaDetails.JenisKelamin = updateAnggota.JenisKelamin
	}

	if updateAnggota.Telepon != "" {
		anggotaDetails.Telepon = updateAnggota.Telepon
	}

	if updateAnggota.Alamat != "" {
		anggotaDetails.Alamat = updateAnggota.Alamat
	}
	db.Save(&anggotaDetails)
	res, _ := json.Marshal(anggotaDetails)
	w.Header().Set("Content-Type", "pkglication/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}
