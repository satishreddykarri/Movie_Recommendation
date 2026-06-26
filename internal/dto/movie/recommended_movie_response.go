package movie

type RecommendedMovieResponse struct {
	ID            uint    `json:"id"`
	Title         string  `json:"title"`
	GenreName     string  `json:"genre_name"`
	AverageRating float64 `json:"average_rating"`
}