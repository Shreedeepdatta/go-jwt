package middleware

import (
	"fmt"
	"net/http"
	"os"
	"time"

	"github.com/Shreedeepdatta/go-jwt/initializers"
	"github.com/Shreedeepdatta/go-jwt/models"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
)

func RequireAuth(ctx *gin.Context) {
	tokenString, err := ctx.Cookie("Auth")
	if err != nil {
		ctx.AbortWithStatus(http.StatusUnauthorized)
	}
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		// Don't forget to validate the alg is what you expect:
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
		}

		// hmacSampleSecret is a []byte containing your secret, e.g. []byte("my_secret_key")
		return []byte(os.Getenv("SECRET_KEY")), nil
	})
	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		if float64(time.Now().Unix()) > claims["exp"].(float64) {
			ctx.AbortWithStatus(http.StatusUnauthorized)
		}
		var user models.User
		initializers.DB.First(&user, claims["sub"])
		if user.ID == 0 {
			ctx.AbortWithStatus(http.StatusUnauthorized)
		}
		ctx.Set("user", user)
		ctx.Next()
	} else {
		ctx.AbortWithStatus(http.StatusUnauthorized)
	}

}
