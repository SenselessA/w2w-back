package models

type FavoritesCreateInput struct {
	UserID  int64
	MovieID string `json:"movieId"`
}

type FavoritesOutput struct {
	UserID  int64  `json:"user_id" db:"user_id"`
	MovieID string `json:"movie_id" db:"movie_id"`
}

type FavoritesByUserOutput struct {
	MovieID   string `json:"movieId" db:"movie_id"`
	Title     string `json:"title" db:"title"`
	TitleOrig string `json:"titleOrig" db:"title_orig"`
	AnimeKind string `json:"animeKind" db:"anime_kind"`
	Status    string `json:"status" db:"all_status"`
	Year      string `json:"year" db:"year"`
	PosterUrl string `json:"posterUrl" db:"poster_url"`
	Rating    string `json:"rating" db:"rating"`
}

type FavoritesUpdateInput struct {
	Email    string `json:"email" db:"email"`
	Username string `json:"username" db:"username"`
	Password string `json:"password" db:"password"`
}

type FavoriteByUserIdInput struct {
	UserID  int64
	MovieID string `json:"movieId"`
}

type FavoriteByUserIdOutput struct {
	UserID  int64  `json:"userId" db:"user_id"`
	MovieID string `json:"movieId" db:"movie_id"`
}
