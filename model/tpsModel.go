package main

import (
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type (
	TPS struct {
		ID uint `json:"id" gorm:"column:id;primaryKey;autoIncrement"`
		NoTPS uint8 `json:"no_tps" gorm:"column:no_tps;"`
		Lokasi string `json:"lokasi" gorm:"column:lokasi;size:100;"`
		Kecamatan string `json:"kecamatan" gorm:"column:kecamatan;size:50;"`
		Nagari string `json:"nagari" gorm:"column:nagari;size:50;"`
		Jorong string `json:"jorong" gorm:"column:jorong;size:50;"`
		JPL uint8 `json:"jpl" gorm:"column:jpl;"`
		IsActive uint8 `json:"is_active" gorm:"column:is_active;"`
	}
)


func (TPS) TableName() string{
	return "tps" //nama table di database
}

func main(){
	cs := "host=localhost user=postgres password=superuser dbname=postgres port=5432 sslmode=disable"
	db, err := gorm.Open(postgres.Open(cs), &gorm.Config{})

	if err != nil{
		panic(err)
	}

	err = db.AutoMigrate(TPS{})
	if err != nil{
		panic(err)
	}

	//create
	//p:= Nagari{IDKecamatan:1, NamaNagari: "Lubuk Gadang"}
	//db.Debug().Create(&p)
	//
	//fmt.Println(p)

}
