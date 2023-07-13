package controllers

import (
	"example/web-service-gin/models"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

type TokenRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}
type TokenResponse struct {
	Email string
	Token string
}

var tokenres TokenResponse

func GenerateToken(context *gin.Context) {
	var request TokenRequest
	var user models.User
	if err := context.ShouldBindJSON(&request); err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		context.Abort()
		return
	}
	// check if email exists and password is correct
	record := models.DB.Where("email = ?", request.Email).First(&user)
	if record.Error != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"error": record.Error.Error()})
		context.Abort()
		return
	}
	credentialError := user.CheckPassword(request.Password)
	if credentialError != nil {
		context.JSON(http.StatusUnauthorized, gin.H{"error": "invalid credentials"})
		context.Abort()
		return
	}
	tokenString, err := GenerateJWT(user.Email)
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		context.Abort()
		return
	}
	tokenres.Email = user.Email
	tokenres.Token = tokenString
	context.JSON(http.StatusOK, gin.H{"token": tokenres})
	//context.JSON(http.StatusOK, gin.H{"token": tokenString})
}
func Ping(context *gin.Context) {
	context.JSON(http.StatusOK, gin.H{"message": "pong"})
}
func GetUser(c *gin.Context) {

	token := c.Request.Header.Get("Authorization")
	if token == "" {
		c.JSON(http.StatusInternalServerError, gin.H{"error": fmt.Sprintf("No Authorization Header Provided")})
		c.Abort()
		return
	}
	claims, err := ValidateToken1(token)
	email := claims.Email

	var user models.User
	if err := models.DB.Where("email = ?", email).First(&user).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": user})
	fmt.Println(err)

	//c.JSON(200, gin.H{"token": claims})
}
