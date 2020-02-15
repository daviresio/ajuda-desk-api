package database

import (
	"fmt"
	"github.com/daviresio/ajuda-desk-api/model"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

var (
	//database-1.crwaolacl8ic.us-east-1.rds.amazonaws.com
	connectionString = fmt.Sprintf("host=%s port=%s user=%s dbname=%s password=%s", "localhost", "5432", "postgres", "ajudadesk", "postgres")
	DB               *gorm.DB
)

func init() {
	var err error
	DB, err = gorm.Open("postgres", connectionString)

	if err != nil {
		println(connectionString)
		println(err)
		panic(err)
	}

	DB.SingularTable(true)

	DB.LogMode(true)

	DB.AutoMigrate(
		&model.Agente{}, &model.Anotacao{}, &model.Artigo{}, &model.CategoriaPublicacao{}, &model.Contato{}, &model.Empresa{},
		&model.Grupo{}, &model.ModeloTicket{}, &model.PastaPublicacao{}, &model.PrioridadeTicket{}, &model.Score{}, &model.StatusTicket{},
		&model.TagTicket{}, &model.Tarefa{}, &model.Ticket{}, &model.TipoEmpresa{}, &model.TipoPlano{}, &model.TipoTicket{}, &model.Usuario{},
	)

}
