package main

import (
	"github.com/worldwidepaniel/ria-course-crud/internal/config"
	"github.com/worldwidepaniel/ria-course-crud/internal/routes"
)

func main() {
	cfg := config.InitializeConfig()
	routes.InitializeRouter(cfg.Server.Port)
}
