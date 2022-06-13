package routes

import (
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/worldwidepaniel/ria-course-crud/internal/handlers"
	"github.com/worldwidepaniel/ria-course-crud/internal/middleware"
)

func InitializeRouter(server_port string) {
	r := gin.Default()
	r.Use(cors.New(cors.Config{
		AllowOrigins: []string{"http://localhost:3000"},
		AllowMethods: []string{"DELETE", "PATCH", "GET", "POST"},
		AllowHeaders: []string{"Origin", "content-type", "x-csrf-token", "token"},
		// AllowHeaders:     []string{"content-type", "token"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
		MaxAge:           12 * time.Hour,
	}))

	r.POST("/login", handlers.Login)
	r.POST("/register", handlers.Register)

	v1 := r.Group("/v1")
	v1.Use(middleware.IsAuthenticated())
	{
		v1.PATCH("/note/:note_id", handlers.ModifyNote)
		v1.DELETE("/note/:note_id", handlers.DeleteNote)
		v1.GET("/note/:note_id", handlers.GetUserNote)
		v1.GET("/countNotes", handlers.CountUserNotes)
		v1.POST("/note", handlers.AddNote)
		v1.GET("/notes", handlers.GetUserNotes)
		v1.GET("/search/:phrase", handlers.SearchNotes)
	}
	r.Run(server_port)
}
