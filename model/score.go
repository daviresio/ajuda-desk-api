package model

type Score struct {
	BaseModel
	Nome string `json:"nome"`
	Empresas []Empresa `gorm:"foreignkey:ScoreId" json:"empresas"`
}
