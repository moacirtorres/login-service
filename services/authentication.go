package services

import (
	"log"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"github.com/hyperyuri/webapi-with-go/models"
)

type Claims struct {
	name string `json:"name"`
	jwt.StandardClaims
}

func AuthMiddleware(p models.Login, c *gin.Context) string {

	var jwtKey = []byte("secretKey")

	expectedPassword := p.Password

	expirationTime := time.Now().Add(5 * time.Minute)

	if expectedPassword != p.Password {
		c.JSON(400, gin.H{
			"mensagem": "Senha esperada diferente da digitada pelo usuário!",
		})
	}

	claims := &Claims{
		name: p.Name,
		StandardClaims: jwt.StandardClaims{
			// In JWT, the expiry time is expressed as unix milliseconds
			ExpiresAt: expirationTime.Unix(),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	tokenString, err := token.SignedString(jwtKey)
	if err != nil {
		log.Fatal("Erro ao assinar o token!")
	}

	return tokenString

}
