package middleware

import (
	"os"
	"time"

	"github.com/golang-jwt/jwt/v4"
	"github.com/joho/godotenv"
	"github.com/whitehackerh/PokeStorage/src/domain/entity"
)

type (
	IAuth interface {
		CreateToken(entity.User) (string, error)
	}
	Auth   struct{}
	Claims struct {
		UserId string
		jwt.StandardClaims
	}
)

var jwtKey = []byte(os.Getenv("JWT_SECRET_KEY"))

func init() {
	err := godotenv.Load("/app/backend/.env")
	if err != nil {
		panic("Error loading .env file")
	}
	jwtKey = []byte(os.Getenv("JWT_SECRET_KEY"))
}

func NewAuth() IAuth {
	return &Auth{}
}

func (a *Auth) CreateToken(user entity.User) (string, error) {
	claims := &Claims{
		UserId: user.Id(),
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(100 * time.Minute).Unix(),
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString(jwtKey)
	if err != nil {
		return "", err
	}
	return tokenString, nil
}
