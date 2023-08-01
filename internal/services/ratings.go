package services

import (
	"github.com/SenselessA/w2w_backend/internal/models"
	"github.com/SenselessA/w2w_backend/internal/repository"
	"strconv"
)

type ServiceRating struct {
	repo *repository.RepoRating
}

func initRating(repo *repository.RepoRating) *ServiceRating {
	return &ServiceRating{repo: repo}
}

func (s *ServiceRating) Create(input models.RatingCreateInput) error {
	rating := models.RatingCreateInput{
		MovieID: input.MovieID,
		UserID:  input.UserID,
		Rating:  input.Rating,
	}

	return s.repo.Create(rating)
}

func (s *ServiceRating) Update(input models.RatingUpdateInput) (*models.RatingUpdateOutput, error) {
	rating := models.RatingUpdateInput{
		MovieID: input.MovieID,
		UserID:  input.UserID,
		Rating:  input.Rating,
	}

	return s.repo.Update(rating)
}

func (s *ServiceRating) Delete(input models.RatingDeleteInput) error {
	rating := models.RatingDeleteInput{
		MovieID: input.MovieID,
		UserID:  input.UserID,
	}

	return s.repo.Delete(rating)
}

func (s *ServiceRating) GetRateByUserId(input models.RatingGetByUserIdInput) (*models.RatingGetByUserIdOutput, error) {
	return s.repo.GetRateByUserId(input)
}

func (s *ServiceRating) GetByUserId(input string) ([]models.RatingOutput, error) {
	id, err := strconv.Atoi(input)
	if err != nil {
		return nil, err
	}
	return s.repo.GetByUserId(id)
}
