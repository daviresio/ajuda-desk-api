package database

import (
	"github.com/daviresio/ajuda-desk-api/model"
)

func SeedAll() {

	for _, data := range seedScore() {
		DB.Create(&data);
	}

	for _, data := range seedTipoPlano() {
		DB.Create(&data);
	}

	for _, data := range seedTipoEmpresa() {
		DB.Create(&data).GetErrors();
	}

}

func seedScore() []model.Score {
	return []model.Score {
		{
			Nome: "Alto risco",
		},
		{
			Nome: "Normal",
		},
		{
			Nome: "Muito bom",
		},
	}
}

func seedTipoPlano() []model.TipoPlano {
	return []model.TipoPlano {
		{
			Nome: "Basico",
		},
		{
			Nome: "Premium",
		},
		{
			Nome: "Enterprise",
		},
	}
}

func seedTipoEmpresa() []model.TipoEmpresa {
	return []model.TipoEmpresa {
		{
			Nome: "Automotivo",
		},
		{
			Nome: "Servicos ao consumidor",
		},
		{
			Nome: "Hotel ou restaurante",
		},
		{
			Nome: "Qualidade de vida",
		},
		{
			Nome: "Imobiliario",
		},
		{
			Nome: "Servicos educacionais",
		},
		{
			Nome: "Servicos familiares",
		},
		{
			Nome: "Midia",
		},
		{
			Nome: "Produtos pessoais",
		},
		{
			Nome: "Outros",
		},
	}
}