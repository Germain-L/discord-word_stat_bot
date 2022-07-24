package models

type Word struct {
	ID     uint   `gorm:"primary_key"`
	Word   string `gorm:"not null"`
	UserId uint   `gorm:"not null"`
	Count  int    `gorm:"not null"`
}
