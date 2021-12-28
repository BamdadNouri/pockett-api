package auth

import (
	"fmt"
	"net/http"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

type AuthService struct{}

func NewAuthService() *AuthService {
	return &AuthService{}
}

func (as *AuthService) Authenticate(c *gin.Context) {

	token := c.GetHeader("token")

	if token == "" {
		c.AbortWithStatusJSON(
			http.StatusUnauthorized,
			gin.H{
				"message": "no token is provided",
			},
		)
		return
	}
	realToken, err := jwt.Parse(token, func(t *jwt.Token) (interface{}, error) {
		if _, ok := t.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method")
		}
		return []byte("SECRET"), nil
	})
	if err != nil || !realToken.Valid || realToken.Claims.(jwt.MapClaims)["unique_id"] == nil {
		c.AbortWithStatusJSON(
			http.StatusUnauthorized,
			gin.H{
				"message": "no token is provided",
			},
		)
		return
	}
	c.Set("userID", realToken.Claims.(jwt.MapClaims)["user_id"].(string))
	c.Next()
}