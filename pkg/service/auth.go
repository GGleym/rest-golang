package service

import (
	"crypto/sha1"
	"fmt"

	todo "github.com/ggleym/rest-golang"
	"github.com/ggleym/rest-golang/pkg/repository"
)

const salt = "fdfdfd323232fdfd"

type AuthService struct {
	repo repository.Authorization
}

func NewAuthService(repo repository.Authorization) *AuthService {
	return &AuthService{repo: repo}
}

func (s *AuthService) CreateUser(user todo.User) (int, error) {
	user.Password = generatePasswordHash(user.Password)

	return s.repo.CreateUser(user)
}

func generatePasswordHash(password string) string {
	hash := sha1.New()
	hash.Write([]byte(password))

	return fmt.Sprintf("%x", hash.Sum([]byte(salt)))
}
