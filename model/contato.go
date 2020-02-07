package model

type Contato struct {
	BaseModel
	Nome string `json:"nome"`
	Cargo string `json:"cargo"`
	Email string `json:"email"`
	Telefone string `json:"telefone"`
	Celular string `json:"celular"`
	Twitter string `json:"twitter"`
	Codigo string `json:"codigo"`
	Foto string `json:"foto"`
	Empresas []Empresa `gorm:"many2many:empresa_contato;" json:"empresas"`
	Tickets []*Ticket `gorm:"foreignkey:ContatoId" json:"tickets"`
}