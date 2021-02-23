package main

import (
"gorm.io/driver/postgres"
"gorm.io/gorm"
)

type (
	AdminRole struct {
		ID uint `json:"id" gorm:"column:id;primaryKey;autoIncrement"`
		RoleAdmin string `json:"role" gorm:"column:role;"`
	}
)


func (AdminRole) TableName() string{
	return "adminrole" //nama table di database
}

func main(){
	cs := "host=localhost user=postgres password=superuser dbname=postgres port=5432 sslmode=disable"
	db, err := gorm.Open(postgres.Open(cs), &gorm.Config{})

	if err != nil{
		panic(err)
	}

	err = db.AutoMigrate(AdminRole{})
	if err != nil{
		panic(err)
	}

	//create
	//p:= Nagari{IDKecamatan:1, NamaNagari: "Lubuk Gadang"}
	//db.Debug().Create(&p)
	//
	//fmt.Println(p)

}

