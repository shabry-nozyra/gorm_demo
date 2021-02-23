package main

import (
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type (
	Admin struct {
		ID uint `json:"id" gorm:"column:id;primaryKey;autoIncrement"`
		Name string `json:"name" gorm:"column:name;size:100;"`
		Email string `json:"email" gorm:"column:email;size:50;"`
		Password string `json:"password" gorm:"column:password;"`
		Image string `json:"image" gorm:"column:image;"`
		RoleAdmin uint `json:"role_id" gorm:"column:role_id;"`
		IsActive uint8 `json:"is_active" gorm:"column:is_active;"`
		TimeCreated uint8 `json:"timecreated" gorm:"column:timecreated;"`
	}
)


func (Admin) TableName() string{
	return "admin" //nama table di database
}

func main(){
	cs := "host=localhost user=postgres password=superuser dbname=postgres port=5432 sslmode=disable"
	db, err := gorm.Open(postgres.Open(cs), &gorm.Config{})

	if err != nil{
		panic(err)
	}

	err = db.AutoMigrate(Admin{})
	if err != nil{
		panic(err)
	}

	//create
	//p:= Nagari{IDKecamatan:1, NamaNagari: "Lubuk Gadang"}
	//db.Debug().Create(&p)
	//
	//fmt.Println(p)

}
