package main

import (
	"github.com/worldwidepaniel/ria-course-crud/internal/config"
	"github.com/worldwidepaniel/ria-course-crud/internal/routes"
)

func main() {
	config.InitializeConfig()
	routes.InitializeRouter(config.AppConfig.Server.Port)
}
