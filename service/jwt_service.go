package service

import (
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"os"
	"time"
)

type JWTService interface {
	GenerateToken(name string, email string) string
	ValidateToken(token string) (*jwt.Token, error)
}
type authCustomClaims struct {
	Name  string `json:"name"`
	Email string `json:"email"`
	jwt.StandardClaims
}

type JWTServiceImpl struct {
	secretKey string
	issuer    string
}

func JWTAuthService() JWTService {
	return &JWTServiceImpl{
		secretKey: getSecretKey(),
		issuer:    "admin",
	}
}

func getSecretKey() string {
	secret := os.Getenv("SECRET")
	if secret == "" {
		secret = "secret"
	}
	return secret
}

func (service *JWTServiceImpl) GenerateToken(name string, email string) string {
	claims := &authCustomClaims{
		name,
		email,
		jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Hour * 48).Unix(),
			Issuer:    service.issuer,
			IssuedAt:  time.Now().Unix(),
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	//encoded string
	t, err := token.SignedString([]byte(service.secretKey))
	if err != nil {
		panic(err)
	}
	return t
}

func (service *JWTServiceImpl) ValidateToken(encodedToken string) (*jwt.Token, error) {
	return jwt.Parse(encodedToken, func(token *jwt.Token) (interface{}, error) {
		if _, isValid := token.Method.(*jwt.SigningMethodHMAC); !isValid {
			return nil, fmt.Errorf("invalid token %s", token.Header["alg"])

		}
		return []byte(service.secretKey), nil
	})

}
