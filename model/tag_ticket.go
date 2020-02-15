package model

type TagTicket struct {
	BaseModel
	Nome string `json:"nome"`
	Tickets []Ticket `gorm:"many2many:ticket_tagticket;" json:"tickets"`
	ModeloTickets []ModeloTicket `gorm:"many2many:modeloticket_tagticket;" json:"modelo_tickets"`
}
