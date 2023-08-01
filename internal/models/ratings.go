package models

type RatingCreateInput struct {
	UserID  int64
	MovieID string `json:"movieId"`
	Rating  uint   `json:"rating" db:"rating"`
}

type RatingOutput struct {
	MovieID   string `json:"movieId" db:"movie_id"`
	Rating    uint   `json:"rating" db:"rating"`
	Title     string `json:"title" db:"title"`
	TitleOrig string `json:"titleOrig" db:"title_orig"`
	AnimeKind string `json:"animeKind" db:"anime_kind"`
	Status    string `json:"status" db:"all_status"`
	Year      string `json:"year" db:"year"`
	PosterUrl string `json:"posterUrl" db:"poster_url"`
}

type RatingUpdateInput struct {
	UserID  int64  `db:"user_id"`
	MovieID string `json:"movieId" db:"movie_id"`
	Rating  uint   `json:"rating" db:"rating"`
}

type RatingUpdateOutput struct {
	UserID  int64  `json:"userId" db:"user_id"`
	MovieID string `json:"movieId" db:"movie_id"`
	Rating  uint   `json:"rating" db:"rating"`
}

type RatingDeleteInput struct {
	UserID  int64  `db:"user_id"`
	MovieID string `json:"movieId" db:"movie_id"`
}

type RatingGetByUserIdInput struct {
	UserID  int64  `json:"userId" db:"user_id"`
	MovieID string `json:"movieId" db:"movie_id"`
}

type RatingGetByUserIdOutput struct {
	Rating uint `json:"rating" db:"rating"`
}
