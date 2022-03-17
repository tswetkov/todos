package service

import (
	"crypto/sha1"
	"fmt"

	"github.com/tswetkov/todos"
	"github.com/tswetkov/todos/pkg/repository"
)

const salt = "1k083vnwd91dvv8120csd87101"

func generatePasswordHash(password string) string {
	hash := sha1.New()
	hash.Write([]byte(password))

	return fmt.Sprintf("%x", hash.Sum([]byte(salt)))
}

type AuthService struct {
	repo repository.Authorization
}

func NewAuthService(repo repository.Authorization) *AuthService {
	return &AuthService{repo}
}

func (s *AuthService) CreateUser(user todos.User) (int, error) {
	user.Password = generatePasswordHash(user.Password)
	return s.repo.CreateUser(user)
}
