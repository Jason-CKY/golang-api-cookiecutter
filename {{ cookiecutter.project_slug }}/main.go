package main

import (
	"flag"
	"fmt"

	"github.com/joho/godotenv"
	"github.com/labstack/echo/v4"
	log "github.com/sirupsen/logrus"
	"github.com/{{ cookiecutter.author }}/{{ cookiecutter.project_slug }}/pkg/core"
	"github.com/{{ cookiecutter.author }}/{{ cookiecutter.project_slug }}/pkg/handlers"
	"github.com/{{ cookiecutter.author }}/{{ cookiecutter.project_slug }}/pkg/utils"
)

func main() {
	// Load environment variables from .env file
	err := godotenv.Load()

	if err != nil {
		log.Infof("Error loading .env file: %v\nUsing environment variables instead...", err)
	}

	flag.StringVar(&core.LogLevel, "log-level", utils.LookupEnvOrString("LOG_LEVEL", core.LogLevel), "Logging level for the server")
	flag.StringVar(&core.DirectusHost, "fpath", utils.LookupEnvOrString("DIRECTUS_HOST", core.DirectusHost), "Hostname for directus server")
	flag.IntVar(&core.WebPort, "port", utils.LookupEnvOrInt("PORT", core.WebPort), "Port for echo web server")
	{% if cookiecutter.use_oauth %}
	flag.StringVar(&core.GithubRedirectUri, "github-redirect-uri", utils.LookupEnvOrString("GITHUB_REDIRECT_URI", core.GithubRedirectUri), "Github redirect uri")
	flag.StringVar(&core.GithubClientID, "github-client-id", utils.LookupEnvOrString("GITHUB_CLIENT_ID", core.GithubClientID), "Github client id")
	flag.StringVar(&core.GithubClientSecret, "github-client-secret", utils.LookupEnvOrString("GITHUB_CLIENT_SECRET", core.GithubClientSecret), "Github client secret")
	{% endif %}
	flag.Parse()

	// setup logrus
	log.SetReportCaller(true)
	log.SetFormatter(&log.TextFormatter{
		FullTimestamp:          true,
		DisableLevelTruncation: true,
	})
	logLevel, _ := log.ParseLevel(core.LogLevel)
	log.SetLevel(logLevel)

	log.Infof("connecting to directus at: %v", core.DirectusHost)

	e := echo.New()
	e.Static("/static", "static")
	e.GET("/", handlers.HomePage)
	{% if cookiecutter.use_oauth %}
	e.GET("/login/github", handlers.LoginRedirect)
	e.GET("/oauth/redirect", handlers.OauthRedirectPage)
	e.GET("/logout", handlers.LogoutRedirect)
	{% endif %}

	e.GET("/htmx", handlers.TasksView)
	e.POST("/htmx/task/empty/:status", handlers.EmptyEditTaskView)
	e.POST("/htmx/task/:id", handlers.EditTaskView)
	e.DELETE("/htmx/task/:id", handlers.DeleteTaskView)
	e.PUT("/htmx/task/:id", handlers.UpdateTaskView)
	e.DELETE("/htmx/task/cancel/:id", handlers.CancelEditTaskView)
	e.POST("/htmx/sort/:status", handlers.SortTaskView)

	e.Logger.Fatal(e.Start(fmt.Sprintf(":%v", core.WebPort)))
}
