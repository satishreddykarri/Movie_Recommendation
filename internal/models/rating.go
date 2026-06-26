package models

import "time"

type Rating struct {
	ID uint `gorm:"column:id;primaryKey;autoIncrement" json:"id"`
	UserID uint `gorm:"column:user_id;not null,uniqueIndex:idx_user_movie" json:"user_id"`
	MovieID uint `gorm:"column:movie_id;not null,uniqueIndex:idx_user_movie" json:"movie_id"`
	Rating int `gorm:"column:rating;type:int;not null" json:"rating"`
	User User `gorm:"foreignKey:UserID;references:ID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;" json:"user"`
	Movie Movie `gorm:"foreignKey:MovieID;references:ID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;" json:"movie"`
	CreatedAt time.Time `gorm:"column:created_at;not null" json:"created_at"`
	UpdatedAt time.Time `gorm:"column:updated_at;not null" json:"updated_at"`
}

func (Rating) TableName() string {
	return "ratings"
}