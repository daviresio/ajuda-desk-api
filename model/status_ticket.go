package model

type StatusTicket struct {
	BaseModel
	Nome string `json:"nome"`
	Tickets []Ticket `gorm:"foreignkey:StatusId" json:"tickets"`
}