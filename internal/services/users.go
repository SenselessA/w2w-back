package services

import (
	"database/sql"
	"github.com/SenselessA/w2w_backend/internal/models"
	"github.com/SenselessA/w2w_backend/internal/repository"
	"github.com/gofiber/fiber/v2"
	"strings"
)

type PasswordHasher interface {
	Hash(password string) (string, error)
}

type ServiceUsers struct {
	repo   *repository.RepoUsers
	hasher PasswordHasher
}

func initUsers(repo *repository.RepoUsers, hasher PasswordHasher) *ServiceUsers {
	return &ServiceUsers{repo: repo, hasher: hasher}
}

func (s *ServiceUsers) Create(input models.UserCreateInput) (int64, error) {
	password, err := s.hasher.Hash(input.Password)
	if err != nil {
		return 0, err
	}

	user := models.UserCreateInput{
		Username: input.Username,
		Email:    input.Email,
		Password: password,
	}

	id, err := s.repo.Create(user)
	if err != nil {
		if strings.Contains(err.Error(), "unique constraint \"users_username_key\"") {
			return 0, fiber.NewError(fiber.StatusBadRequest, "Такое имя пользователя уже занято")
		}

		if strings.Contains(err.Error(), "unique constraint \"users_email_key\"") {
			return 0, fiber.NewError(fiber.StatusBadRequest, "Такой email уже зарегистрирован")
		}

		return 0, err
	}

	return id, nil
}

func (s *ServiceUsers) Login(input models.UserLoginInput) (*models.UserOutput, error) {
	password, err := s.hasher.Hash(input.Password)
	if err != nil {
		return nil, err
	}

	user, err := s.repo.FindByEmail(input.Email)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, fiber.NewError(fiber.StatusNotFound, "Пользователь с таким email не найден")
		}
		return nil, err
	}

	if user.Password != password {
		return nil, fiber.NewError(fiber.StatusUnauthorized, "Неверный пароль")
	}

	return user, nil
}

func (s *ServiceUsers) GetUserById(input string) (*models.UserInfoOutput, error) {
	user, err := s.repo.FindById(input)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, fiber.NewError(fiber.StatusNotFound, "Пользователь с таким id не найден")
		}
		return nil, err
	}

	return user, nil
}
