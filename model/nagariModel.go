package main

import (
	"fmt"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type (
	Nagari struct {
	ID uint `json:"id_nagari" gorm:"column:id_nagari;primaryKey;autoIncrement"`
	IDKecamatan uint `json:"id_kecamatan" gorm:"column:id_kecamatan;"`
	NamaNagari string `json:"nama_nagari" gorm:"column:nama_nagari;size:100;"`
		}
)


func (Nagari) TableName() string{
	return "nagari" //nama table di database
}

func main(){
	cs := "host=localhost user=postgres password=superuser dbname=postgres port=5432 sslmode=disable"
	db, err := gorm.Open(postgres.Open(cs), &gorm.Config{})

	if err != nil{
		panic(err)
	}

	err = db.AutoMigrate(Nagari{})
	if err != nil{
		panic(err)
	}

	//create
	p:= Nagari{IDKecamatan:1, NamaNagari: "Lubuk Gadang"}
	db.Debug().Create(&p)

	fmt.Println(p)

}