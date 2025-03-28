package service

import (
	"errors"

	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"

	"AvitoTest/internal/constants"
	"AvitoTest/internal/model/dto"
	"AvitoTest/internal/model/entity"
	"AvitoTest/internal/repository"
	"AvitoTest/internal/util"
)

type UserService struct {
	userRepo *repository.UserRepository
}

func NewUserService(userRepo *repository.UserRepository) *UserService {
	return &UserService{userRepo: userRepo}
}

func (s *UserService) AuthUser(authRequest dto.AuthRequest) (string, error) {

	user, err := s.userRepo.FindUserByUsername(authRequest.Username)

	if errors.Is(err, gorm.ErrRecordNotFound) {
		if user, err = s.registerUser(authRequest); err != nil {
			return "", err
		}
		return util.GenerateToken(user)
	}

	if err == nil {
		if bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(authRequest.Password)) != nil {
			return "", errors.New("invalid password")
		}
		return util.GenerateToken(user)
	}

	return constants.EmptyString, errors.New("auth user error")
}

func (s *UserService) registerUser(authRequest dto.AuthRequest) (*entity.User, error) {
	cryptPass, cryptErr := cryptPassword(authRequest.Password)
	if cryptErr != nil {
		return nil, errors.New("password crypt error")
	}

	user := entity.User{
		Username: authRequest.Username,
		Password: cryptPass,
		Balance:  constants.StartBalance,
	}

	result := s.userRepo.CreateUser(&user)

	if result.Error != nil {
		return nil, result.Error
	}

	return &user, nil
}

func cryptPassword(password string) (string, error) {
	genPass, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return constants.EmptyString, err
	}
	return string(genPass), err
}
