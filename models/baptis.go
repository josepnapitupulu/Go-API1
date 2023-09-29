package models

import (
	"github.com/jinzhu/gorm"
	"github.com/josepnapitupulu/API_Baptis/config"
	)

var db *gorm.DB

type Baptis struct {
	gorm.Model
	NIK string `gorm:""json:"nik"`
	NamaAnak string `json:"nama_anak"`
	TanggalBaptis string `json:"tanggal_Baptis"`
	TanggalLahir string `json:"tanggal_lahir"`
	NamaAyah string `json:"nama_ayah"`
	NamaIbu string `json:"nama_ibu"`
	Alamat string `json:"alamat"`
}

func init() {
	config.Connect()
	db = config.GetDB()
	db.AutoMigrate(&Baptis{})
}

func (b *Baptis) CreateBaptis() *Baptis {
	db.NewRecord(b)
	db.Create(&b)
	return b
}

func GetAllBaptiss() []Baptis {
	var Baptiss []Baptis
	db.Find(&Baptiss)
	return Baptiss
}

func GetBaptisbyId(nik int64) (*Baptis, *gorm.DB) {
	var getBaptis Baptis
	db := db.Where("NIK=?", nik).Find(&getBaptis)
	return &getBaptis, db
}

func DeleteBaptis(nik int64) Baptis {
	var baptis Baptis
	db.Where("NIK=?", nik).Delete(baptis)
	return baptis 
}