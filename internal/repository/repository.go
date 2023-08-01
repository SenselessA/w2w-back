package repository

import (
	"github.com/jmoiron/sqlx"
)

type Repository struct {
	Users     *RepoUsers
	Favorites *RepoFavorites
	Rating    *RepoRating
	Movies    *RepoMovies
}

func New(db *sqlx.DB) *Repository {
	return &Repository{
		Users:     initUsers(db),
		Favorites: initFavorites(db),
		Rating:    initRating(db),
		Movies:    initMovies(db),
	}
}
