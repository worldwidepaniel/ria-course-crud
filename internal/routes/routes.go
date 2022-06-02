package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/worldwidepaniel/ria-course-crud/internal/handlers"
)

func InitializeRouter(server_port string) {
	r := gin.Default()
	r.GET("/ping", handlers.Pong)
	r.Run(server_port)
}
