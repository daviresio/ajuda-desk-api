package model

type ModeloTicket struct {
	Titulo string `json:"titulo"`
	DescricaoModelo string `json:"descricao_modelo"`
	//caso seja nullo ou vazio tornar disponivel para todos
	DisponivelPara *[]Grupo `gorm:"many2many:modeloticket_grupo;" json:"disponivel_para"`
	Assunto string `json:"assunto"`
	Tipo *TipoTicket `gorm:"foreignkey:TipoId" json:"tipo"`
	TipoId uint64 `json:"tipo_id"`
	Status *StatusTicket `gorm:"foreignkey:StatusId" json:"status"`
	StatusId uint64 `json:"status_id"`
	Prioridade *PrioridadeTicket `gorm:"foreignkey:PrioridadeId" json:"prioridade"`
	PrioridadeId uint64 `json:"prioridade_id"`
	Agente *Agente `gorm:"foreignkey:AgenteId" json:"agente"`
	AgenteId uint64 `json:"agente_id"`
	Descricao string `json:"descricao"`
	Tags []TagTicket `gorm:"many2many:modeloticket_tagticket;" json:"tags"`
}
