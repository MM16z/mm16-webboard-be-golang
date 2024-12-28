package port

import "github.com/MM16z/mm16-webboard-be-golang/internal/core/dto"

type AuthRepository interface {
	Register(req dto.RegisterRequest) (registerResponse dto.RegisterResponse, err error)
}

type AuthService interface {
	Register(req dto.RegisterRequest) (registerResponse dto.RegisterResponse, err error)
}
