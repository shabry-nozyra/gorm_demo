package main

import (
"gorm.io/driver/postgres"
"gorm.io/gorm"
)

type (
	Petugas struct {
		ID uint `json:"id" gorm:"column:id;primaryKey;autoIncrement"`
		NoTPS uint8 `json:"no_tps" gorm:"column:no_tps;"`
		Kecamatan string `json:"kecamatan" gorm:"column:kecamatan;size:100;"`
		Nagari string `json:"nagari" gorm:"column:nagari;size:100;"`
		Jorong string `json:"jorong" gorm:"column:jorong;size:100;"`
		Foto uint8 `json:"foto" gorm:"column:foto;"`
		NoHp uint8 `json:"no_hp" gorm:"column:no_hp;"`
		Email string `json:"email" gorm:"column:email;"`
		Username string `json:"username" gorm:"column:username;size:50;"`
		NamaLengkap string `json:"nama_lengkap" gorm:"column:nama_lengkap;size:50;"`
		Password string `json:"password" gorm:"column:password;"`
		RolePetugas uint8 `json:"role_petugas" gorm:"column:role_petugas;"`
		IsActive uint8 `json:"is_active" gorm:"column:is_active;"`
		TimeKirim uint8 `json:"timekirim" gorm:"column:timekirim;"`
	}
)


func (Petugas) TableName() string{
	return "petugas" //nama table di database
}

func main(){
	cs := "host=localhost user=postgres password=superuser dbname=postgres port=5432 sslmode=disable"
	db, err := gorm.Open(postgres.Open(cs), &gorm.Config{})

	if err != nil{
		panic(err)
	}

	err = db.AutoMigrate(Petugas{})
	if err != nil{
		panic(err)
	}

	//create
	//p:= Nagari{IDKecamatan:1, NamaNagari: "Lubuk Gadang"}
	//db.Debug().Create(&p)
	//
	//fmt.Println(p)

}
