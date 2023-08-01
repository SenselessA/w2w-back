package repository

import (
	"github.com/SenselessA/w2w_backend/internal/models"
	"github.com/jmoiron/sqlx"
)

type RepoUsers struct {
	db *sqlx.DB
}

func initUsers(db *sqlx.DB) *RepoUsers {
	return &RepoUsers{db: db}
}

func (r *RepoUsers) FindByEmail(email string) (*models.UserOutput, error) {
	var user models.UserOutput

	err := r.db.QueryRowx(`
		SELECT id, email, username, password, registration_time FROM users WHERE email = $1
	`, email).StructScan(&user)
	if err != nil {
		return nil, err
	}

	return &user, nil
}

func (r *RepoUsers) FindById(id string) (*models.UserInfoOutput, error) {
	var user models.UserInfoOutput

	err := r.db.QueryRowx(`
		SELECT id, email, username, registration_time FROM users WHERE id = $1
	`, id).StructScan(&user)
	if err != nil {
		return nil, err
	}

	return &user, nil
}

func (r *RepoUsers) Create(input models.UserCreateInput) (int64, error) {
	var id int64

	err := r.db.QueryRow(`
	INSERT INTO users (username, email, password) VALUES ($1, $2, $3) RETURNING id
	`, input.Username, input.Email, input.Password).Scan(&id)
	if err != nil {
		return 0, err
	}

	return id, nil
}
