package rating

type Response struct {
	ID        uint   `json:"id"`
	UserID    uint   `json:"user_id"`
	UserName  string `json:"user_name"`
	MovieID   uint   `json:"movie_id"`
	MovieName string `json:"movie_name"`
	Rating    int    `json:"rating"`
}