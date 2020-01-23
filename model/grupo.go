package model

type Grupo struct {
	BaseModel
	Nome string `gorm:"not null" json:"nome"`
}