package model

type Tarefa struct {
	BaseModel
	Nome string `json:"nome"`
	Concluido bool `json:"concluido"`
	Empresa *Empresa `gorm:"foreignkey:EmpresaId" json:"empresa"`
	EmpresaId uint64 `json:"empresa_id"`
}
