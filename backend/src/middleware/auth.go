package middleware

import (
	"net/http"
	"os"
	"strings"
	"time"

	"github.com/gin-gonic/gin"

	"github.com/golang-jwt/jwt/v4"
	"github.com/joho/godotenv"
	"github.com/whitehackerh/PokeStorage/src/domain/entity"
	"github.com/whitehackerh/PokeStorage/src/model"
	"github.com/whitehackerh/PokeStorage/src/util"
)

type (
	IAuth interface {
		CreateToken(entity.User) (string, error)
		ParseToken() gin.HandlerFunc
		IsBlacklisted(string) (bool, error)
		CreateBlacklist(string, Claims) error
	}
	IJwtBlacklistRepository interface {
		FindByToken(string) (model.JwtBlacklist, error)
		Create(model.JwtBlacklist) error
	}
	Auth struct {
		repo IJwtBlacklistRepository
	}
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

func NewAuth(jwtBlacklistRepository IJwtBlacklistRepository) IAuth {
	return &Auth{
		repo: jwtBlacklistRepository,
	}
}

func (a *Auth) CreateToken(user entity.User) (string, error) {
	claims := &Claims{
		UserId: user.Id(),
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(153722867 * time.Minute).Unix(),
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString(jwtKey)
	if err != nil {
		return "", err
	}
	return tokenString, nil
}

func (a *Auth) ParseToken() gin.HandlerFunc {
	return func(c *gin.Context) {
		handleUnauthorized := func(errMessage string) {
			c.JSON(http.StatusUnauthorized, gin.H{"error": errMessage})
			c.Abort()
		}

		authorization := c.GetHeader("Authorization")
		if authorization == "" {
			handleUnauthorized("Missing token")
			return
		}
		tokenString := strings.Split(authorization, "Bearer ")[1]
		claims := &Claims{}
		token, err := jwt.ParseWithClaims(
			tokenString,
			claims,
			func(token *jwt.Token) (interface{}, error) {
				return jwtKey, nil
			},
		)
		if err != nil || !token.Valid {
			handleUnauthorized("Invalid token")
			return
		}
		if isBlacklist, err := a.IsBlacklisted(tokenString); isBlacklist || err != nil {
			handleUnauthorized("Token is blacklisted")
			return
		}
		c.Set("token", tokenString)
		c.Set("claims", claims)
		c.Next()
	}
}

func (a *Auth) IsBlacklisted(token string) (bool, error) {
	if _, err := a.repo.FindByToken(token); err != nil {
		return false, nil
	} else {
		return true, err
	}
}

func (a *Auth) CreateBlacklist(token string, claims Claims) error {
	return a.repo.Create(model.JwtBlacklist{Id: util.NewUUID(), Token: token, ExpiresAt: time.Unix(claims.ExpiresAt, 0)})
}
