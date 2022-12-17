package model

type BoxModel struct {
	Id          uint64 `gorm:"primary_key,column:id"`
	Name        string `gorm:"column:name"`
	Color       string `gorm:"column:color"`
	Tag         string `gorm:"column:tag"`
	Description string `gorm:"column:description"`
}
