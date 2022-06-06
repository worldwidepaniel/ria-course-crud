package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/worldwidepaniel/ria-course-crud/internal/handlers"
	"github.com/worldwidepaniel/ria-course-crud/internal/middleware"
)

func InitializeRouter(server_port string) {
	r := gin.Default()
	r.POST("/login", handlers.Login)
	r.POST("/register", handlers.Register)

	v1 := r.Group("/v1")
	v1.Use(middleware.IsAuthenticated())
	{
		v1.PATCH("/note/:note_id", handlers.ModifyNote)
		v1.DELETE("/note/:note_id", handlers.DeleteNote)
		v1.POST("/note/", handlers.AddNote)
		v1.GET("/note/:note_id", handlers.GetNote)
	}
	r.Run(server_port)
}
