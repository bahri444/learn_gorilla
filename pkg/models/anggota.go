package models

import (
	"learn_gorilla/pkg/config"

	"github.com/jinzhu/gorm"
)

var db *gorm.DB

type Anggota struct {
	gorm.Model
	Nim          string `gorm:""json:"nim"`
	Nama         string `json:"nama"`
	JenisKelamin string `json: "jenis_kelamin`
	Alamat       string `json:"alamat"`
	Telepon      string `json:"telepon"`
}

func init() {
	config.Connect()
	db = config.GetDB()
	db.AutoMigrate(&Anggota{})
}

func (a *Anggota) CreateAnggota() *Anggota {
	db.NewRecord(a)
	db.Create(&a)
	return a
}

func GetAllAnggota() []Anggota {
	var Anggotas []Anggota
	db.Find(&Anggotas)
	return Anggotas
}

func GetAnggotaById(Id int64) (*Anggota, *gorm.DB) {
	var getAnggota Anggota
	db := db.Where("ID=?", Id).Find(&getAnggota)
	return &getAnggota, db
}
func DeleteAnggota(ID int64) Anggota {
	var anggota Anggota
	db.Where("ID=?", ID).Delete(anggota)
	return anggota
}
