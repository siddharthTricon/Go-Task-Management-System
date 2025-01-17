package services

import (
	"go/token"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/siddharthTricon/go-task-management-sysytem/utils"
)

type JWTService interface{
	GenerateToken(userID uint, role string) (string, error)
	ValidateToken(token string) (*jwt.Token, error)
}

type jwtCustomClaims struct{
	UserID uint `json:"user-id"`
	Role string `json:"role"`
	jwt.RegisteredClaims
}

type jwtService struct{
	secretKey string
}

func NewJWTService() JWTService{
	return &jwtService{secretKey: utils.GetEnv("JWT_SECRET")}
}

func (s *jwtService) GenerateToken(userID uint, role string) (string, error){
	claims := &jwtCustomClaims{
		UserID: userID,
		Role: role,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt : jwtNewNumericDate(time.Now().Add(72 * time.Hour)),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte(s.secretKey))
}

func (s *jwtService) ValidateToken(token string) (*jwt.Token, error){
	return jwt.Pasre(token, func(t *jwt.Token) (interface{}, error){
		return []byte(s.secretKey), nil
	})
}

