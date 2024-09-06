package user

import (
	"time"

	"golang.org/x/crypto/bcrypt"
)

type Service interface {
	RegisterUser(input RegisterUserInput) (User, error)
}

type service struct {
	repository Repository
}

func NewService(repository Repository) *service {
	return &service{repository}
}

func (s *service) RegisterUser(input RegisterUserInput) (User, error) {
	user := User{}

	user.Name = input.Name
	user.Email = input.Email
	user.Occupation = input.Occupation

	password_hash, err := bcrypt.GenerateFromPassword([]byte(input.Password), bcrypt.MinCost)
	if err != nil {
		return user, err
	}

	user.Password_hash = string(password_hash)
	user.Role = "user"
	user.Created_at = time.Now()
	user.Updated_at = time.Now()

	NewUser,err := s.repository.Save(user)
	if err != nil {
		return NewUser, err
	}
	
	return NewUser, nil
}