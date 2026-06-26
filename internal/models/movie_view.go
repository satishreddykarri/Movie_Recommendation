package models
type MovieWithGenre struct {
	ID          uint
	Title       string
	Description string
	ReleaseYear int
	GenreID     uint
	GenreName   string
}