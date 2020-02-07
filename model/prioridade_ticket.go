package model

type PrioridadeTicket struct {
	BaseModel
	Nome string `json:"nome"`
	Tickets []Ticket `gorm:"foreignkey:PrioridadeId" json:"tickets"`
}