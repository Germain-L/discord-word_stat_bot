package models

type Users struct {
	ID   uint   `gorm:"primary_key"`
	Name string `gorm:"not null"`
}
