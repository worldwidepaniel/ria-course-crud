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
	err = db.DeleteNote(objectID)
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

}

func GetNote(c *gin.Context) {

}
