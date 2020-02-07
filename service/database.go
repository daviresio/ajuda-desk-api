package service

import (
	"fmt"
	"github.com/daviresio/ajuda-desk-api/model"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

var (
	connectionString = fmt.Sprintf("host=%s port=%s user=%s dbname=%s password=%s", "database-1.crwaolacl8ic.us-east-1.rds.amazonaws.com", "5432", "postgres", "ajudadesk", "postgres")
	db *gorm.DB
)

func init() {
	var err error
	db, err = gorm.Open("postgres", connectionString)

	if err != nil {
		println(connectionString)
		println(err)
		panic(err)
	}

	db.SingularTable(true)

	db.LogMode(true)

	db.AutoMigrate(&model.Grupo{}, &model.Contato{}, &model.Empresa{}, &model.Score{}, &model.Ticket{})

}
