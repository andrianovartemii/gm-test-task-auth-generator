package tokenManager

import (
	"crypto/rand"
	"encoding/base64"
	"errors"
	"github.com/golang-jwt/jwt/v5"
	"time"
)

type Manager struct {
	Secret   string
	Lifetime time.Duration
}

func GenerateRandomSecret(length int) (string, error) {
	numBytes := length / 4 * 3
	if length%4 != 0 {
		numBytes += 3
	}
	bytes := make([]byte, numBytes)
	_, err := rand.Read(bytes)
	if err != nil {
		return "", err
	}
	secret := base64.URLEncoding.EncodeToString(bytes)
	if len(secret) > length {
		secret = secret[:length]
	}
	return secret, nil
}

func (manager *Manager) GenerateToken(login string) (string, error) {
	expirationTime := time.Now().Add(manager.Lifetime)
	claims := &Claims{
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(expirationTime),
		},
		Login: login,
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte(manager.Secret))
}

func (manager *Manager) ValidateToken(tokenString string) error {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		return []byte(manager.Secret), nil
	})
	if err != nil {
		if err == jwt.ErrTokenExpired {
			return errors.New("token has expired")
		}
		return err
	}
	if !token.Valid {
		return errors.New("token is invalid")
	}
	return nil
}
