package handlers

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/worldwidepaniel/ria-course-crud/internal/db"
	"github.com/worldwidepaniel/ria-course-crud/internal/utils"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func AddNote(c *gin.Context) {
	var requestBody db.Note
	user := utils.UserEmailFromJWT(c.Request.Header["Token"][0])
	user_id := db.GetUser(user)
	if err := c.BindJSON(&requestBody); err != nil {
		c.AbortWithStatusJSON(
			http.StatusBadRequest,
			gin.H{
				"error": err,
			})
		return
	}
	requestBody.UID = user_id.UID
	requestBody.Note_ID = primitive.NewObjectID()
	created_note_id := db.AddNote(requestBody)
	if created_note_id == "" {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"error": created_note_id,
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"success": fmt.Sprintf("created note with id: %s", created_note_id),
	})

}

func DeleteNote(c *gin.Context) {
	noteID := c.Param("note_id")
	objectID, err := primitive.ObjectIDFromHex(noteID)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"error": "invalid note id",
		})
		return
	}
	email := utils.UserEmailFromJWT(c.Request.Header["Token"][0])
	user := db.GetUser(email)
	err = db.DeleteNote(objectID, user.UID)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"error": "error while deleting note",
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"success": fmt.Sprintf("deleted note of id: %s", objectID),
	})

}

func ModifyNote(c *gin.Context) {
	noteID := c.Param("note_id")
	objectID, err := primitive.ObjectIDFromHex(noteID)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"error": "invalid note id",
		})
		return
	}
	var requestBody db.Note
	if err := c.BindJSON(&requestBody); err != nil {
		c.AbortWithStatusJSON(
			http.StatusBadRequest,
			gin.H{
				"error": err,
			})
		return
	}
	email := utils.UserEmailFromJWT(c.Request.Header["Token"][0])
	user := db.GetUser(email)
	requestBody.Note_ID = objectID
	requestBody.UID = user.UID
	if err = db.ModifyNote(requestBody); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"error": err,
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"success": "modified note",
	})
}

func GetUserNotes(c *gin.Context) {
	notesLimit := c.DefaultQuery("limit", "10")
	notesOffset := c.DefaultQuery("offset", "0")

	email := utils.UserEmailFromJWT(c.Request.Header["Token"][0])
	user := db.GetUser(email)
	result, err := db.GetUserNotes(notesLimit, notesOffset, user.UID)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"error": err,
		})
		return
	}

	c.JSON(http.StatusOK, result)

}

func SearchNotes(c *gin.Context) {
	phrase := c.Param("phrase")
	email := utils.UserEmailFromJWT(c.Request.Header["Token"][0])
	user := db.GetUser(email)
	hits := db.SearchPhrase(phrase, user.UID)

	c.JSON(http.StatusOK, hits)
}
