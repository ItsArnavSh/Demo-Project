package userservice

import (
	"context"
	entity "demo/application/internal"
	"demo/application/internal/repository"

	"github.com/google/uuid"
	"go.uber.org/zap"
)

type UserService struct {
	database repository.Repository
	logger   *zap.Logger
}

func NewService(ctx context.Context, logger *zap.Logger, database repository.Repository) *UserService {
	return &UserService{database, logger}
}

func (u *UserService) RegisterUser(ctx context.Context, staff entity.Staff) (string, error) {
	var err error
	staff.StaffId = uuid.NewString()
	staff.HashedPassword, err = HashPassword(staff.HashedPassword)
	if err != nil {
		u.logger.Error("Could not hash password")
		return "", err
	}
	jwt, err := GenerateJWT(staff.StaffId, staff.Type)
	if err != nil {
		u.logger.Error("Could not generate JWT")
		return "", err
	}
	err = u.database.RegisterUser(ctx, staff)
	if err != nil {
		u.logger.Error("Failed to Register User")
		return "", err
	}
	return jwt, nil
}
func (u *UserService) LoginUser(ctx context.Context, id string, password string, Type string) (string, error) {
	// Retrieve the stored hashed password from DB
	storedHash, err := u.database.GetPassword(ctx, id)
	if err != nil {
		u.logger.Error("Failed to get stored password hash")
		return "", err
	}

	// Compare input password with stored hash
	if err := ComparePassword(password, storedHash); err != nil {
		u.logger.Error("Invalid password")
		return "", err
	}

	// Generate JWT after successful login
	jwt, err := GenerateJWT(id, Type)
	if err != nil {
		u.logger.Error("Failed to generate JWT")
		return "", err
	}

	return jwt, nil
}

func (u *UserService) CheckPassword(ctx context.Context, userID, password string) (bool, error) {
	hashedPass, err := u.database.GetPassword(ctx, userID)
	if err != nil {
		u.logger.Error("Could not fetch password from DB")
		return false, err
	}

	if err := ComparePassword(password, hashedPass); err != nil {
		return false, nil // password mismatch
	}

	return true, nil // match
}
