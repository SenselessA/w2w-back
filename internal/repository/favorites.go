package repository

import (
	"github.com/SenselessA/w2w_backend/internal/models"
	"github.com/jmoiron/sqlx"
)

type RepoFavorites struct {
	db *sqlx.DB
}

func initFavorites(db *sqlx.DB) *RepoFavorites {
	return &RepoFavorites{db: db}
}

func (r *RepoFavorites) Create(input models.FavoritesCreateInput) error {
	_, err := r.db.Exec(`
	INSERT INTO user_favorite_movies (user_id, movie_id) VALUES ($1, $2)
	`, input.UserID, input.MovieID)
	if err != nil {
		return err
	}

	return nil
}

func (r *RepoFavorites) Delete(input models.FavoritesCreateInput) error {
	_, err := r.db.Exec(`
	DELETE FROM user_favorite_movies WHERE user_id = $1 AND movie_id = $2
	`, input.UserID, input.MovieID)
	if err != nil {
		return err
	}

	return nil
}

func (r *RepoFavorites) GetByUserId(input int) ([]models.FavoritesByUserOutput, error) {

	var favorites []models.FavoritesByUserOutput

	err := r.db.Select(&favorites, `
		SELECT ufm.movie_id, m.title, m.title_orig, md.anime_kind, md.all_status,
		       md.year, md.poster_url, md.kinopoisk_rating AS rating
		FROM user_favorite_movies as ufm, movies as m, material_data as md
		         WHERE ufm.user_id = $1 AND ufm.movie_id = m.id AND ufm.movie_id = md.movie_id
	`, input)
	if err != nil {
		return nil, err
	}

	return favorites, nil
}

func (r *RepoFavorites) GetFavoriteByUserId(input models.FavoriteByUserIdInput) (*models.FavoriteByUserIdOutput, error) {
	var favorite models.FavoriteByUserIdOutput

	err := r.db.QueryRowx(`
		SELECT ufm.movie_id, ufm.user_id
		FROM user_favorite_movies as ufm
		         WHERE ufm.user_id = $1 AND ufm.movie_id = $2
	`, input.UserID, input.MovieID).StructScan(&favorite)
	if err != nil {
		return nil, err
	}

	return &favorite, nil
}
