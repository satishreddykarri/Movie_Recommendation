package models

type RecommendedMovie struct {
	ID            uint    `gorm:"column:id"`
	Title         string  `gorm:"column:title"`
	GenreName     string  `gorm:"column:genre_name"`
	AverageRating float64 `gorm:"column:average_rating"`
}