package model

type TipoEmpresa struct {
	BaseModel
	Nome string `json:"nome"`
	Empresas []Empresa `gorm:"foreignkey:TipoEmpresaId" json:"empresas"`
}
