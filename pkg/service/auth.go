package service

import (
	"crypto/sha1"
	"errors"
	"fmt"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/tswetkov/todos"
	"github.com/tswetkov/todos/pkg/repository"
)

type tokenClaims struct {
	jwt.StandardClaims
	UserId int `json:"user_id"`
}

const (
	salt      = "1k083vnwd91dvv8120csd87101"
	signinKey = "asdfasdf1231fw12vw"
	tokenTTL  = 12 * time.Hour
)

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

func (s *AuthService) GenerateToken(username string, password string) (string, error) {
	user, err := s.repo.GetUser(username, generatePasswordHash(password))
	if err != nil {
		return "", err
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, &tokenClaims{
		jwt.StandardClaims{
			ExpiresAt: time.Now().Add(tokenTTL).Unix(),
			IssuedAt:  time.Now().Unix(),
		}, user.Id,
	})

	return token.SignedString([]byte(signinKey))
}

func (s *AuthService) ParseToken(authToken string) (int, error) {
	token, err := jwt.ParseWithClaims(authToken, &tokenClaims{}, func(t *jwt.Token) (interface{}, error) {
		if _, ok := t.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errors.New("invalid signing method")
		}

		return []byte(signinKey), nil
	})

	if err != nil {
		return 0, err
	}

	clames, ok := token.Claims.(*tokenClaims)
	if !ok {
		return 0, errors.New("token clames are not of type *tokenClames")
	}

	return clames.UserId, nil
}
