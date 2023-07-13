package controllers

import (
	"example/web-service-gin/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

func RegisterUser(context *gin.Context) {
	var user models.User
	if err := context.ShouldBindXML(&user); err != nil {
		context.XML(http.StatusBadRequest, gin.H{"error": err.Error()})
		context.Abort()
		return
	}
	if err := user.HashPassword(user.Password); err != nil {
		context.XML(http.StatusInternalServerError, gin.H{"error": err.Error()})
		context.Abort()
		return
	}
	record := models.DB.Create(&user)
	if record.Error != nil {
		context.XML(http.StatusInternalServerError, gin.H{"error": record.Error.Error()})
		context.Abort()
		return
	}
	context.XML(http.StatusCreated, gin.H{"name": user.Name, "email": user.Email, "mobilenumber": user.MobileNumber, "password": user.Password})
}
