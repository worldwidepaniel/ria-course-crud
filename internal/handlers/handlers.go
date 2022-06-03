package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/worldwidepaniel/ria-course-crud/internal/db"
)

func Pong(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "pong",
	})
}

func Login(c *gin.Context) {
	var requestBody EmailRequestBody
	if err := c.BindJSON(&requestBody); err != nil {
		c.AbortWithStatusJSON(
			http.StatusBadRequest,
			gin.H{
				"error": "malformed body request",
			})
		return
	}
	user := db.GetUser(requestBody.Email)
	if user == nil {
		c.JSON(http.StatusOK, gin.H{})
		return
	}
	c.JSON(http.StatusOK, user)

}
