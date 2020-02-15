package model

type Anotacao struct {
	BaseModel
	Titulo string `json:"titulo"`
	Conteudo string `json:"conteudo"`
	Empresa *Empresa `gorm:"foreignkey:EmpresaId" json:"empresa"`
	EmpresaId uint64 `json:"empresa_id"`
}
