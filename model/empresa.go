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
	Score *Score `gorm:"foreignkey:ScoreId" json:"score"`
	ScoreId uint64 `json:"score_id"`
	TipoPlano *TipoPlano `gorm:"foreignkey:TipoPlanoId" json:"tipo_plano"`
	TipoPlanoId uint64 `json:"tipo_plano_id"`
	TipoEmpresa *TipoEmpresa `gorm:"foreignkey:TipoEmpresaId" json:"tipo_empresa"`
	TipoEmpresaId uint64 `json:"tipo_empresa_id"`
	DataRenovacao *time.Time `json:"data_renovacao"`
	Dominios pq.StringArray `gorm:"type:varchar(100)[]" json:"dominios"`
	Contatos []*Contato `gorm:"many2many:empresa_contato;" json:"contatos"`
	Tarefas []Tarefa `gorm:"foreignkey:EmpresaId" json:"tarefas"`
	Anotacoes []Anotacao `gorm:"foreignkey:EmpresaId" json:"anotacoes"`
}

