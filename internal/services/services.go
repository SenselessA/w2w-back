package services

import "github.com/SenselessA/w2w_backend/internal/repository"

type Services struct {
	Users     *ServiceUsers
	Favorites *ServiceFavorites
	Ratings   *ServiceRating
	Kodik     *ServiceKodik
	Movies    *ServiceMovies
}

func New(repo *repository.Repository, hasher PasswordHasher) *Services {
	return &Services{
		Users:     initUsers(repo.Users, hasher),
		Favorites: initFavorites(repo.Favorites),
		Ratings:   initRating(repo.Rating),
		Kodik:     initKodik(repo.Movies),
		Movies:    initMoviesService(repo.Movies),
	}
}
