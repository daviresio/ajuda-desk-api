package model

import (
	"github.com/lib/pq"
	"time"
)

type Empresa struct {
	BaseModel
	Nome string `json:"nome"`
	Foto *string `json:"foto"`
	Descricao *string `json:"descricao"`
	Anotacao *string `json:"anotacao"`
	Score int64 `json:"score"`
	TipoPlano int64 `json:"tipo_plano"`
	TipoEmpresa int64 `json:"tipo_empresa"`
	DataRenovacao *time.Time `json:"data_renovacao"`
	Dominios pq.StringArray `gorm:"type:varchar(100)[]" json:"dominios"`
	Contatos []Contato `gorm:"many2many:empresa_contato;" json:"contatos"`
	Tarefas []Tarefa `gorm:"foreignkey:EmpresaId" json:"tarefas"`
}