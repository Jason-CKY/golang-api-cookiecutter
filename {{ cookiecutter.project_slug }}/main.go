package main

import (
	"flag"
	"fmt"

	"github.com/labstack/echo/v4"
	log "github.com/sirupsen/logrus"
	"github.com/{{ cookiecutter.author }}/{{ cookiecutter.project_slug }}/pkg/core"
	"github.com/{{ cookiecutter.author }}/{{ cookiecutter.project_slug }}/pkg/handlers"
	"github.com/{{ cookiecutter.author }}/{{ cookiecutter.project_slug }}/pkg/utils"
)

func main() {
	flag.StringVar(&core.DirectusHost, "fpath", utils.LookupEnvOrString("DIRECTUS_HOST", core.DirectusHost), "Hostname for directus server")
	flag.IntVar(&core.WebPort, "port", utils.LookupEnvOrInt("PORT", core.WebPort), "Port for echo web server")

	flag.Parse()

	// setup logrus
	log.SetFormatter(&log.TextFormatter{
		FullTimestamp:          true,
		DisableLevelTruncation: true,
	})

	log.Infof("connecting to directus at: %v", core.DirectusHost)

	e := echo.New()
	e.Static("/static", "static")
	e.GET("/", handlers.HomePage)
	e.GET("/htmx", handlers.TasksView)
	e.POST("/htmx/task/empty/:status", handlers.EmptyEditTaskView)
	e.POST("/htmx/task/:id", handlers.EditTaskView)
	e.DELETE("/htmx/task/:id", handlers.DeleteTaskView)
	e.PUT("/htmx/task/:id", handlers.UpdateTaskView)
	e.DELETE("/htmx/task/cancel/:id", handlers.CancelEditTaskView)
	e.POST("/htmx/sort/:status", handlers.SortTaskView)

	e.Logger.Fatal(e.Start(fmt.Sprintf(":%v", core.WebPort)))
}
