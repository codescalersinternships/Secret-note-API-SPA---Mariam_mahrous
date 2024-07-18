package middleware

import (
	"errors"
	"fmt"
	"net/http"
	"strings"
	"time"

	database "github.com/codescalersinternships/Secret-note-API-SPA-Mariam_mahrous/database"
	model "github.com/codescalersinternships/Secret-note-API-SPA-Mariam_mahrous/models"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	"golang.org/x/crypto/bcrypt"
)

type MW struct {
	TokenKey string
	DB database.DB
}

func (m *MW) RequireAuth(c *gin.Context) {
	tokenString := strings.Split(c.Request.Header["Authorization"][0], " ")[1]

	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}

		return []byte(m.TokenKey), nil
	})

	if err != nil {
		c.AbortWithStatus(http.StatusUnauthorized)
		return
	}

	if claims, ok := token.Claims.(jwt.MapClaims); ok {
		if float64(time.Now().Unix()) > claims["exp"].(float64) {
			c.AbortWithStatus(http.StatusUnauthorized)
			return
		}
		authorized, user := m.DB.VerifyUser(claims["email"])
		if authorized {
			c.Set("id", user.ID)
		}
		c.Next()
	} else {
		c.AbortWithStatus(http.StatusUnauthorized)
		return
	}
}

func (m *MW) GenerateToken(user model.User) (string, int, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"id":    user.ID,
		"email": user.Email,
		"exp":   time.Now().Add(time.Hour * 24 * 30).Unix(),
	})

	tokenString, err := token.SignedString([]byte(m.TokenKey))

	if err != nil {
		return "", http.StatusConflict, errors.New("couldn't create Token")
	}
	return tokenString, 0, nil
}

func (m *MW) HashPassword(password string) (string ,int , error) {
	hashedPassword, err := model.HashPassword(password)
	if err != nil {
		return  "",http.StatusConflict, errors.New("error hashing the password")
	}
	return hashedPassword , 200 , nil
}

func (m *MW) CompareHashAndPassword(currentPassword , enteredPassword string) (int , error) {
	err := bcrypt.CompareHashAndPassword([]byte(currentPassword), []byte(enteredPassword))
	if err != nil {
		return  http.StatusConflict, errors.New("invalid Password or email")
	}
	return 200 , nil
}