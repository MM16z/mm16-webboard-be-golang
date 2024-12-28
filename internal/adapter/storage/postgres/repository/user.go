package repository

import (
	"github.com/MM16z/mm16-webboard-be-golang/internal/adapter/storage/postgres"
	"github.com/MM16z/mm16-webboard-be-golang/internal/core/dto"
)

type UserRepository struct {
	db *postgres.DB
}

func NewUserRepository(db *postgres.DB) *UserRepository {
	return &UserRepository{db: db}
}

// TODO: Bcrypt
func (ur *UserRepository) Register(req dto.RegisterRequest) (dto.RegisterResponse, error) {
	tx := ur.db.Begin()
	if tx.Error != nil {
		return dto.RegisterResponse{}, tx.Error
	}

	user := dto.RegisterRequest{
		Username: req.Username,
		Password: req.Password,
		Email:    req.Email,
	}

	result := tx.Table("users").Create(&user)
	if result.Error != nil {
		tx.Rollback()
		return dto.RegisterResponse{}, result.Error
	}

	if err := tx.Commit().Error; err != nil {
		tx.Rollback()
		return dto.RegisterResponse{}, err
	}

	return dto.RegisterResponse{
		Code:    200,
		Message: "User created successfully",
		Status:  "success",
		Data:    user,
	}, nil
}
