package models

import "time"

type User struct {
	ID uint `gorm:"column:id;primaryKey;autoIncrement" json:"id"`
	Name string `gorm:"column:name;type:varchar(100);not null" json:"name"`
	Email string `gorm:"column:email;type:varchar(255);unique;not null" json:"email"`
	CreatedAt time.Time `gorm:"column:created_at;not null" json:"created_at"`
	UpdatedAt time.Time `gorm:"column:updated_at;not null" json:"updated_at"`
}

func (User) TableName() string {
	return "users"
}