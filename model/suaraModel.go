package main

import (
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type (
	Suara struct {
		ID uint `json:"id" gorm:"column:id;primaryKey;autoIncrement"`
		IdTps string 	`json:"role" gorm:"column:role;"`
		Suara1 uint8 	`json:"suara1" gorm:"column:suara1;"`
		Suara2 uint8 	`json:"suara2" gorm:"column:suara2;"`
		Suara3 uint8 	`json:"suara3" gorm:"column:suara3;"`
		SuaraSah uint8 	`json:"suarasah" gorm:"column:suarasah;"`
		JPL uint8 		`json:"jpl" gorm:"column:jpl;"`
		SuaraTidakSah uint8 	`json:"suaratidaksah" gorm:"column:suaratidaksah;"`
		JumlahGolput uint 		`json:"jumlahgolput" gorm:"column:jumlahgolput;"`
		FotoC1 string 			`json:"fotoc1" gorm:"column:fotoc1;"`
		Status uint8			`json:"status" gorm:"column:status;"`
		WaktuInput uint8		`json:"waktuinput" gorm:"column:waktuinput;"`
		DataKe uint8 			`json:"datake" gorm:"column:datake;"`


	}
)


func (Suara) TableName() string{
	return "suara" //nama table di database
}

func main(){
	cs := "host=localhost user=postgres password=superuser dbname=postgres port=5432 sslmode=disable"
	db, err := gorm.Open(postgres.Open(cs), &gorm.Config{})

	if err != nil{
		panic(err)
	}

	err = db.AutoMigrate(Suara{})
	if err != nil{
		panic(err)
	}

	//create
	//p:= Nagari{IDKecamatan:1, NamaNagari: "Lubuk Gadang"}
	//db.Debug().Create(&p)
	//
	//fmt.Println(p)

}

