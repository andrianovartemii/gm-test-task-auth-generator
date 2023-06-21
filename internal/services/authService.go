package services

import "gm-test-task-auth-generator/internal/tokenManager"

type AuthService struct {
	tokenManager *tokenManager.Manager
}

func NewAuthService(tokenManager *tokenManager.Manager) *AuthService {
	return &AuthService{
		tokenManager: tokenManager,
	}
}

func (service *AuthService) GenerateToken(login string) (string, error) {
	return service.tokenManager.GenerateToken(login)
}

func (service *AuthService) ValidateToken(tokenString string) error {
	return service.tokenManager.ValidateToken(tokenString)
}
