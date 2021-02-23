package main

import (
"gorm.io/driver/postgres"
"gorm.io/gorm"
)

type (
	Kecurangan struct {
		ID uint 			`json:"id" gorm:"column:id;primaryKey;autoIncrement"`
		Detail string 		`json:"halaman" gorm:"column:halaman;size:100"`
		WaktuKejadian string `json:"waktu" gorm:"column:waktu;"`
		Jam uint 			`json:"jam" gorm:"column:jam;"`
		Tempat string 		`json:"tempat" gorm:"column:tempat;"`
		Bukti1 string 		`json:"bukti1" gorm:"column:bukti1;"`
		Bukti2 string 		`json:"bukti2" gorm:"column:bukti2;"`
		Bukti3 string 		`json:"bukti3" gorm:"column:bukti3;"`
	}
)


func (Kecurangan) TableName() string{
	return "kecurangan" //nama table di database
}

func main(){
	cs := "host=localhost user=postgres password=superuser dbname=postgres port=5432 sslmode=disable"
	db, err := gorm.Open(postgres.Open(cs), &gorm.Config{})

	if err != nil{
		panic(err)
	}

	err = db.AutoMigrate(Kecurangan{})
	if err != nil{
		panic(err)
	}

	//create
	//p:= Nagari{IDKecamatan:1, NamaNagari: "Lubuk Gadang"}
	//db.Debug().Create(&p)
	//
	//fmt.Println(p)

}



