package service

import (
	"AvitoTest/internal/repository"
	"errors"
	"golang.org/x/crypto/bcrypt"
	"testing"

	"AvitoTest/internal/model/dto"
	"AvitoTest/internal/model/entity"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"gorm.io/gorm"
)

func TestUserService_AuthUser(t *testing.T) {
	mockRepo := new(repository.MockUserRepository)
	service := NewUserService(mockRepo)

	tests := []struct {
		name        string
		authRequest dto.AuthRequest
		mockUser    *entity.User
		mockError   error
		wantErr     bool
	}{
		{
			name:        "Регистрация нового пользователя",
			authRequest: dto.AuthRequest{Username: "newUser", Password: "password123"},
			mockUser:    nil,
			mockError:   gorm.ErrRecordNotFound,
			wantErr:     false,
		},
		{
			name:        "Успешная аутентификация",
			authRequest: dto.AuthRequest{Username: "existingUser", Password: "password123"},
			mockUser:    &entity.User{Username: "existingUser", Password: hashPassword("password123")},
			mockError:   nil,
			wantErr:     false,
		},
		{
			name:        "Ошибка: Неверный пароль",
			authRequest: dto.AuthRequest{Username: "existingUser", Password: "wrongpass"},
			mockUser:    &entity.User{Username: "existingUser", Password: hashPassword("password123")},
			mockError:   nil,
			wantErr:     true,
		},
		{
			name:        "Ошибка: Проблема с БД",
			authRequest: dto.AuthRequest{Username: "errorUser", Password: "password123"},
			mockUser:    nil,
			mockError:   errors.New("DB error"),
			wantErr:     true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mockRepo.On("FindUserByUsername", tt.authRequest.Username, mock.Anything).Return(tt.mockUser, tt.mockError)

			if errors.Is(tt.mockError, gorm.ErrRecordNotFound) {
				mockRepo.On("CreateUser", mock.Anything).Return(&gorm.DB{})
			}

			token, err := service.AuthUser(tt.authRequest)

			if tt.wantErr {
				assert.Error(t, err)
				assert.Empty(t, token)
			} else {
				assert.NoError(t, err)
				assert.NotEmpty(t, token)
			}

			mockRepo.AssertExpectations(t)
		})
	}
}

func TestUserService_registerUser(t *testing.T) {
	mockRepo := new(repository.MockUserRepository)
	service := NewUserService(mockRepo)

	tests := []struct {
		name        string
		authRequest dto.AuthRequest
		mockCrypt   func()
		mockDBError error
		wantErr     bool
	}{
		{
			name:        "Успешная регистрация",
			authRequest: dto.AuthRequest{Username: "newUser", Password: "password123"},
			mockDBError: nil,
			wantErr:     false,
		},
		{
			name:        "Ошибка при сохранении в БД",
			authRequest: dto.AuthRequest{Username: "newUser", Password: "password123"},
			mockDBError: errors.New("DB error"),
			wantErr:     true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			if tt.mockDBError != nil {
				mockRepo.On("CreateUser", mock.Anything).Return(&gorm.DB{Error: tt.mockDBError}).Once()
			} else {
				mockRepo.On("CreateUser", mock.Anything).Return(&gorm.DB{Error: nil}).Once()
			}

			user, err := service.registerUser(tt.authRequest)
			if (err != nil) != tt.wantErr {
				t.Errorf("registerUser() error = %v, wantErr %v", err, tt.wantErr)
			}

			if tt.wantErr && user != nil {
				t.Errorf("registerUser() should return nil user, got: %+v", user)
			}

			mockRepo.AssertExpectations(t)
		})
	}
}

func hashPassword(password string) string {
	hashed, _ := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	return string(hashed)
}
