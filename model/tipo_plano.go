package model

type TipoPlano struct {
	BaseModel
	Nome string `json:"nome"`
	Empresas []Empresa `gorm:"foreignkey:TipoPlanoId" json:"empresas"`
}
