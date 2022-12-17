package model

type TaskModel struct {
	Id          uint64 `gorm:"primary_key,column:id"`
	Name        string `gorm:"column:name"`
	Status      string `gorm:"column:status"`
	Color       string `gorm:"column:color"`
	Tag         string `gorm:"column:tag"`
	Group       string `gorm:"column:Group"`
	Description string `gorm:"column:description"`
}
