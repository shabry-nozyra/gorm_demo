package main

import (
"fmt"
"gorm.io/driver/postgres"
"gorm.io/gorm"
)

type Paslon struct {
	ID uint `json:"id" gorm:"column:id;primaryKey;autoIncrement"`
	Bupati string `json:"bupati" gorm:"column:bupati; size:200"`
	Wakil string `json:"wakil" gorm:"column:wakil; size:200"`
	NoUrut uint8 `json:"no_urut" gorm:"column:no_urut;"`
}

func (Paslon) TableName() string{
	return "paslon" //nama table di database
}

func main(){
	cs := "host=localhost user=postgres password=superuser dbname=postgres port=5432 sslmode=disable"
	db, err := gorm.Open(postgres.Open(cs), &gorm.Config{})

	if err != nil{
		panic(err)
	}

	err = db.AutoMigrate(Paslon{})
	if err != nil{
		panic(err)
	}

	//create
	p:= Paslon{Bupati:"Erwin Ali", Wakil: "Marwan Efendi", NoUrut: 3}
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
