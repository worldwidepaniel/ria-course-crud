package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/worldwidepaniel/ria-course-crud/internal/db"
	"github.com/worldwidepaniel/ria-course-crud/internal/utils"
	"golang.org/x/crypto/bcrypt"
)

func Pong(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "pong",
	})
}

func Login(c *gin.Context) {
	var requestBody LoginRequestBody
	if err := c.BindJSON(&requestBody); err != nil {
		c.AbortWithStatusJSON(
			http.StatusBadRequest,
			gin.H{
				"error": "malformed body request",
			})
		return
	}
	user := db.GetUser(requestBody.Email)
	if user == (db.User{}) {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "no user with given email",
		})
		return
	}

	if err := bcrypt.CompareHashAndPassword([]byte(user.PasswordHash), []byte(requestBody.Password)); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "wrong password",
		})
		return
	}
	token, err := utils.GenerateJWT(user.Email)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "error while generating token",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"token": token,
	})

}

func Register(c *gin.Context) {
	var requestBody RegisterRequestBody
	if err := c.BindJSON(&requestBody); err != nil {
		c.AbortWithStatusJSON(
			http.StatusBadRequest,
			gin.H{
				"error": "malformed body request",
			})
		return
	}
	if user := db.GetUser(requestBody.Email); user != (db.User{}) {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"error": "user with this email already exists",
		})
		return
	}

	hash, err := utils.GetPasswordHash([]byte(requestBody.Password))
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"error": "error while hashing password",
		})
	}

	newUser := db.User{
		Email:        requestBody.Email,
		Name:         requestBody.Name,
		PasswordHash: hash,
	}

	if userCreation := db.CreateUser(newUser); userCreation != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"error": userCreation,
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"success": "new user created succesfuly",
	})

}
