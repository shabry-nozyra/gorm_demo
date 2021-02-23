package main

import (
	"fmt"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type Jorong struct {
	ID uint `json:"id" gorm:"column:id;primaryKey;autoIncrement"`
	NamaJorong string `json:"nama_jorong" gorm:"column:nama_jorong;"`
	NamaNagari string `json:"nama_nagari" gorm:"column:nama_nagari;"`
	NamaKecamatan string `json:"nama_kecamatan" gorm:"column:nama_kecamatan;"`
	IsActive uint8 `json:"is_active" gorm:"column:is_active;"`
}

func (Jorong) TableName() string{
	return "data_jorong" //nama table di database
}

func main(){
	cs := "host=localhost user=postgres password=superuser dbname=postgres port=5432 sslmode=disable"
	db, err := gorm.Open(postgres.Open(cs), &gorm.Config{})

	if err != nil{
		panic(err)
	}

	err = db.AutoMigrate(Jorong{})
	if err != nil{
		panic(err)
	}

	//create
	p:= Jorong{NamaJorong:"pakan rabaa", NamaNagari: "pakan rabaa", NamaKecamatan: "koto parik gadang diateh", IsActive: 1}
	db.Debug().Create(&p)


	//Select ALl
	//var ps []Paslon
	//db.Find(&ps)
	//fmt.Println(ps)

	//select one
	//p:= Paslon{}
	//db.Where("id = ?",1).First(&p)
	//fmt.Println(p)

	//update
	//dapatkan data terlebih dahulu

	//p:= Paslon{}
	//db.Where("id = ?",1).First(&p)
	//
	//p.Bupati = "Khairunnas"
	//p.Wakil = "Yulian Efi"
	//p.NoUrut = 1
	//
	//db.Updates(&p)

	fmt.Println(p)

}
