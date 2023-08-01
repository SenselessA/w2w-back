package services

import (
	"github.com/SenselessA/w2w_backend/internal/models"
	"github.com/SenselessA/w2w_backend/internal/repository"
	"strconv"
)

type ServiceFavorites struct {
	repo *repository.RepoFavorites
}

func initFavorites(repo *repository.RepoFavorites) *ServiceFavorites {
	return &ServiceFavorites{repo: repo}
}

func (s *ServiceFavorites) Create(input models.FavoritesCreateInput) error {
	favorite := models.FavoritesCreateInput{
		MovieID: input.MovieID,
		UserID:  input.UserID,
	}

	return s.repo.Create(favorite)
}

func (s *ServiceFavorites) Delete(input models.FavoritesCreateInput) error {
	favorite := models.FavoritesCreateInput{
		MovieID: input.MovieID,
		UserID:  input.UserID,
	}

	return s.repo.Delete(favorite)
}

func (s *ServiceFavorites) GetByUserId(input string) ([]models.FavoritesByUserOutput, error) {
	id, err := strconv.Atoi(input)
	if err != nil {
		return nil, err
	}
	return s.repo.GetByUserId(id)
}

func (s *ServiceFavorites) GetFavoriteByUserId(input models.FavoriteByUserIdInput) (*models.FavoriteByUserIdOutput, error) {
	return s.repo.GetFavoriteByUserId(input)
}
