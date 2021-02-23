package main

import (
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type (
	Pesan struct {
		ID uint 		`json:"id" gorm:"column:id;primaryKey;autoIncrement"`
		Name string 	`json:"name" gorm:"column:name;size:40"`
		Phone string 	`json:"phone" gorm:"column:phone; size:15"`
		Email string 	`json:"email" gorm:"column:email;size:40"`
		Perihal string `json:"perihal" gorm:"column:perihal;"`
		Message string `json:"message" gorm:"column:message;"`
		TimeCreated uint8 `json:"timecreate" gorm:"column:timecreate;"`
	}
)


func (Pesan) TableName() string{
	return "pesan" //nama table di database
}

func main(){
	cs := "host=localhost user=postgres password=superuser dbname=postgres port=5432 sslmode=disable"
	db, err := gorm.Open(postgres.Open(cs), &gorm.Config{})

	if err != nil{
		panic(err)
	}

	err = db.AutoMigrate(Pesan{})
	if err != nil{
		panic(err)
	}

	//create
	//p:= Nagari{IDKecamatan:1, NamaNagari: "Lubuk Gadang"}
	//db.Debug().Create(&p)
	//
	//fmt.Println(p)

}

