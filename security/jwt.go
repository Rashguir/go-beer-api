package security

import (
	"time"
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

type userClaims struct {
	Role string `json:"role"`
	jwt.StandardClaims
}

var signingKey = []byte("SecretKeyThatShouldBeLongEnough")

// GenerateToken returns a JSON response containing a new JWT
func GenerateToken(c *gin.Context) {
	claims := userClaims{
		"owner",
		jwt.StandardClaims{
			ExpiresAt: time.Now().Add(1 * time.Hour).Unix(),
			Issuer: "api",
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	ss, err := token.SignedString(signingKey)
	if err != nil {
		c.JSON(500, gin.H{"status": "error"})
		return
	}
	c.JSON(200, gin.H{"status": "success", "token": ss})
}
