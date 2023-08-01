package services

import (
	"github.com/SenselessA/w2w_backend/internal/models"
	"github.com/SenselessA/w2w_backend/internal/repository"
)

type ServiceMovies struct {
	repo *repository.RepoMovies
}

func initMoviesService(repo *repository.RepoMovies) *ServiceMovies {
	return &ServiceMovies{repo: repo}
}

func (s *ServiceMovies) GetByMovieId(input string) (*models.MovieDataByMovieIdOutput, error) {
	return s.repo.GetByMovieId(input)
}
