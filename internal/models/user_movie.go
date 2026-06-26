package models

type UserMovie struct {
	MovieID   uint   `gorm:"column:movie_id"`
	MovieName string `gorm:"column:movie_name"`
	Rating    int    `gorm:"column:rating"`
}