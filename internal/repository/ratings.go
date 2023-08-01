package repository

import (
	"github.com/SenselessA/w2w_backend/internal/models"
	"github.com/jmoiron/sqlx"
)

type RepoRating struct {
	db *sqlx.DB
}

func initRating(db *sqlx.DB) *RepoRating {
	return &RepoRating{db: db}
}

func (r *RepoRating) Create(input models.RatingCreateInput) error {
	_, err := r.db.Exec(`
		INSERT INTO user_movie_rating (user_id, movie_id, rating) VALUES ($1, $2, $3)
		ON CONFLICT (user_id, movie_id) DO UPDATE SET
	  	(user_id, movie_id, rating) = ($1, $2, $3)
	`, input.UserID, input.MovieID, input.Rating)
	if err != nil {
		return err
	}

	return nil
}

func (r *RepoRating) Update(input models.RatingUpdateInput) (*models.RatingUpdateOutput, error) {
	updatedMovieRating := models.RatingUpdateOutput{}

	err := r.db.QueryRowx(`
		UPDATE user_movie_rating SET (user_id, movie_id, rating) = ($1, $2, $3)
		WHERE user_id = $1 AND movie_id = $2
		RETURNING user_id, movie_id, rating;
	`, input.UserID, input.MovieID, input.Rating).StructScan(&updatedMovieRating)
	if err != nil {
		return nil, err
	}

	return &updatedMovieRating, nil
}

func (r *RepoRating) Delete(input models.RatingDeleteInput) error {
	_, err := r.db.Exec(`
		DELETE FROM user_movie_rating WHERE user_id = $1 AND movie_id = $2
	`, input.UserID, input.MovieID)
	if err != nil {
		return err
	}

	return nil
}

func (r *RepoRating) GetRateByUserId(input models.RatingGetByUserIdInput) (*models.RatingGetByUserIdOutput, error) {
	var rating models.RatingGetByUserIdOutput

	err := r.db.QueryRowx(`
		SELECT umr.rating
		FROM user_movie_rating as umr
		         WHERE umr.user_id = $1 AND umr.movie_id = $2
	`, input.UserID, input.MovieID).StructScan(&rating)
	if err != nil {
		println("GetRateByUserId: ", err.Error())
		return nil, err
	}

	return &rating, nil
}

func (r *RepoRating) GetByUserId(input int) ([]models.RatingOutput, error) {
	var ratings []models.RatingOutput

	err := r.db.Select(&ratings, `
		SELECT umr.movie_id, umr.rating, m.title, m.title_orig, md.anime_kind, md.all_status,
		       md.year, md.poster_url
		FROM user_movie_rating as umr, movies as m, material_data as md
		         WHERE umr.user_id = $1 AND umr.movie_id = m.id AND umr.movie_id = md.movie_id
	`, input)
	if err != nil {
		println("GetByUserId: ", err.Error())
		return nil, err
	}

	return ratings, nil
}
