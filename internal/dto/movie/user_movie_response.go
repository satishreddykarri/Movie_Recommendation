package movie

type UserMovieResponse struct {
	MovieID   uint   `json:"movie_id"`
	MovieName string `json:"movie_name"`
	Rating    int    `json:"rating"`
}