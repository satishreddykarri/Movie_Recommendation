package models

import "time"

type Movie struct {
	ID uint `gorm:"column:id;primaryKey;autoIncrement" json:"id"`
	Title string `gorm:"column:title;type:varchar(255);not null" json:"title"`
	Description string `gorm:"column:description;type:text" json:"description"`
	ReleaseYear int `gorm:"column:release_year;not null" json:"release_year"`
	GenreID uint `gorm:"column:genre_id;not null" json:"genre_id"`
	Genre Genre `gorm:"foreignKey:GenreID;references:ID" json:"genre"`
	CreatedAt time.Time `gorm:"column:created_at;not null" json:"created_at"`
	UpdatedAt time.Time `gorm:"column:updated_at;not null" json:"updated_at"`
}