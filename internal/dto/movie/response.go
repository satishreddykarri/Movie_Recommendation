package movie

type Response struct {
	ID          uint   `json:"id"`
	Title       string `json:"title"`
	Description string `json:"description"`
	ReleaseYear int    `json:"release_year"`
	GenreID     uint   `json:"genre_id"`
	GenreName	string `json:"genre_name"`
}