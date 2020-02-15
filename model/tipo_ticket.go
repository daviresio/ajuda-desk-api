package model

type TipoTicket struct {
	BaseModel
	Nome string `json:"nome"`
	Tickets []Ticket `gorm:"foreignkey:TipoId" json:"tickets"`
	ModeloTickets []ModeloTicket `gorm:"foreignkey:TipoId" json:"modelo_tickets"`
}