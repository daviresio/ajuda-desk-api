package model

import "github.com/lib/pq"

type Ticket struct {
	BaseModel
	Contato *Contato `gorm:"foreignkey:ContatoId" json:"contato"`
	ContatoId uint64 `json:"contato_id"`
	Cc *pq.StringArray `gorm:"type:varchar(100)[]" json:"cc"`
	Assunto string `json:"assunto"`
	Descricao string `json:"descricao"`
	Tags []TagTicket `gorm:"many2many:ticket_tagticket;" json:"tags"`
	Tipo *TipoTicket `gorm:"foreignkey:TipoId" json:"tipo"`
	TipoId uint64 `json:"tipo_id"`
	Status *StatusTicket `gorm:"foreignkey:StatusId" json:"status"`
	StatusId uint64 `json:"status_id"`
	Prioridade *PrioridadeTicket `gorm:"foreignkey:PrioridadeId" json:"prioridade"`
	PrioridadeId uint64 `json:"prioridade_id"`

}
