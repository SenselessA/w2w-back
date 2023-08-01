package models

import "time"

type UserCreateInput struct {
	Email    string `json:"email" db:"email"`
	Username string `json:"username" db:"username"`
	Password string `json:"password" db:"password"`
}

type UserLoginInput struct {
	Email    string `json:"email" db:"email"`
	Password string `json:"password" db:"password"`
}

type UserUpdateInput struct {
	Email    string `json:"email" db:"email"`
	Username string `json:"username" db:"username"`
	Password string `json:"password" db:"password"`
}

type UserOutput struct {
	Id               int64     `json:"id" db:"id"`
	Username         string    `json:"username" db:"username"`
	Email            string    `json:"email" db:"email"`
	Password         string    `json:"password" db:"password"`
	RegistrationTime time.Time `json:"registration_time" db:"registration_time"`
}

type UserInfoOutput struct {
	Id               int64     `json:"id" db:"id"`
	Username         string    `json:"username" db:"username"`
	Email            string    `json:"email" db:"email"`
	RegistrationTime time.Time `json:"registrationTime" db:"registration_time"`
}

type JwtClaims struct {
	Id    int64     `json:"id" db:"id"`
	Email string    `json:"email" db:"email"`
	Admin bool      `json:"admin" db:"admin"`
	Exp   time.Time `json:"exp" db:"exp"`
}
