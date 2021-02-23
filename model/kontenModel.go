package main

import (
"gorm.io/driver/postgres"
"gorm.io/gorm"
)

type (
	Konten struct {
		ID uint 			`json:"id" gorm:"column:id;primaryKey;autoIncrement"`
		Halaman string 		`json:"halaman" gorm:"column:halaman;size:100"`
		Konten string 		`json:"konten" gorm:"column:konten; size:100"`
		IsiKonten string 	`json:"isikonten" gorm:"column:isikonten;"`
		URL string 			`json:"url" gorm:"column:url;"`
		IsActive uint8 		`json:"isactive" gorm:"column:isactive;"`
	}
)


func (Konten) TableName() string{
	return "konten" //nama table di database
}

func main(){
	cs := "host=localhost user=postgres password=superuser dbname=postgres port=5432 sslmode=disable"
	db, err := gorm.Open(postgres.Open(cs), &gorm.Config{})

	if err != nil{
		panic(err)
	}

	err = db.AutoMigrate(Konten{})
	if err != nil{
		panic(err)
	}

	//create
	//p:= Nagari{IDKecamatan:1, NamaNagari: "Lubuk Gadang"}
	//db.Debug().Create(&p)
	//
	//fmt.Println(p)

}


