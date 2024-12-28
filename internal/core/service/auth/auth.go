package service

import (
	"github.com/MM16z/mm16-webboard-be-golang/internal/core/dto"
	port "github.com/MM16z/mm16-webboard-be-golang/internal/core/port/auth"
)

type AuthService struct {
	repo port.AuthRepository
}

func NewAuthService(repo port.AuthRepository) *AuthService {
	return &AuthService{repo: repo}
}

func (s *AuthService) Register(req dto.RegisterRequest) (registerResponse dto.RegisterResponse, err error) {
	registerResponse, err = s.repo.Register(req)
	if err != nil {
		return dto.RegisterResponse{}, err
	}
	return registerResponse, nil
}
