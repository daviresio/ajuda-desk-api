package model

type Agente struct {
	BaseModel

	ModeloTickets []ModeloTicket `gorm:"foreignkey:AgenteId" json:"modelo_tickets"`
}
