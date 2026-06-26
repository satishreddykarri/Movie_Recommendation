package models

type RatingWithDetails struct {
	ID        uint
	UserID    uint
	UserName  string
	MovieID   uint
	MovieName string
	Rating    int
}