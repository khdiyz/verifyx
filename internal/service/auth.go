package service

import (
	"errors"
	"time"
	"verifyx/config"

	"github.com/golang-jwt/jwt/v4"
	"google.golang.org/grpc/codes"
)

type jwtCustomClaim struct {
	jwt.StandardClaims
	Username string `json:"username"`
	RoleName string `json:"role"`
}

type AuthService struct {
	cfg *config.Config
}

func NewAuthService(cfg *config.Config) *AuthService {
	return &AuthService{cfg: cfg}
}

func (s *AuthService) GenerateToken(role, username string) (string, error) {
	claims := jwt.MapClaims{
		"username": username,
		"role":     role,
		"exp":      time.Now().Add(time.Duration(s.cfg.JWTAccessExpirationHours * int(time.Hour))).Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte(s.cfg.JWTSecret))
}

func (s *AuthService) ParseToken(token string) (*jwtCustomClaim, error) {
	jwtToken, err := jwt.ParseWithClaims(token, &jwtCustomClaim{}, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, serviceError(errors.New("invalid signing method"), codes.InvalidArgument)
		}

		return []byte(config.GetConfig().JWTSecret), nil
	})
	if err != nil {
		return nil, err
	}

	claims, ok := jwtToken.Claims.(*jwtCustomClaim)
	if !ok {
		return nil, serviceError(errors.New("token claims are not of type *jwtCustomClaim"), codes.InvalidArgument)
	}

	return claims, nil
}
