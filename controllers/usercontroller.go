package controllers

import (
	"fmt"
	"net/http"
	"os"
	"time"

	"github.com/Shreedeepdatta/go-jwt/initializers"
	"github.com/Shreedeepdatta/go-jwt/models"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	"golang.org/x/crypto/bcrypt"
)

func SignUp(ctx *gin.Context) {
	var body struct {
		Email    string
		Password string
	}
	if ctx.Bind(&body) != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": "Failed to read body",
		})
		return
	}
	hash, err := bcrypt.GenerateFromPassword([]byte(body.Password), 10)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": "failed hashing",
		})
		return
	}
	user := models.User{Email: body.Email, Password: string(hash)}
	result := initializers.DB.Create(&user)
	if result.Error != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"message": "couldn't sign up",
		})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{
		"message": "User Created",
	})
}
func Login(ctx *gin.Context) {
	var body struct {
		Email    string
		Password string
	}
	if ctx.Bind(&body) != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": "Failed to read body",
		})
		return
	}
	var user models.User
	initializers.DB.First(&user, "email=?", body.Email)
	if user.ID == 0 {
		ctx.JSON(http.StatusNotFound, gin.H{
			"message": "Sign Up First",
		})
		return
	}
	err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(body.Password))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"message": "invalid password",
		})
		return
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"sub": user.ID,
		"exp": time.Now().Add(time.Hour * 24 * 30).Unix(),
	})
	tokenString, err := token.SignedString([]byte(os.Getenv("SECRET_KEY")))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"message": "tokengeneration failed",
		})
		fmt.Println(err)
		return
	}
	ctx.SetSameSite(http.SameSiteLaxMode)
	ctx.SetCookie("Auth", tokenString, 3600*24*30, "", "", false, true)
	ctx.JSON(http.StatusOK, gin.H{})
}
func Validate(ctx *gin.Context){
	user,_:=ctx.Get("user")
	ctx.JSON(http.StatusOK, gin.H{
		"message":user,
	})
}