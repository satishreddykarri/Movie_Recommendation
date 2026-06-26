package models

import "time"

type Genre struct{
	ID uint `gorm:"column:id;primaryKey;autoIncrement" json:"id"`
	Name string `gorm:"column:name;type:varchar(255);not null" json:"name"`
	CreatedAt time.Time `gorm:"column:created_at;autoCreateTime" json:"created_at"`
	UpdatedAt time.Time `gorm:"column:updated_at;autoUpdateTime" json:"updated_at"`
}

func (Genre) TableName() string {
	return "genres"
}