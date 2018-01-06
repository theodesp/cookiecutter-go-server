package main

import (
	"net/http"
	"log"
	"github.com/{{ cookiecutter.github_username }}/{{ cookiecutter.project_slug }}/routes"
	"github.com/{{ cookiecutter.github_username }}/{{ cookiecutter.project_slug }}/utils"
	"github.com/{{ cookiecutter.github_username }}/{{ cookiecutter.project_slug }}/config"
)

func main()  {
	mux := http.NewServeMux()

	utils.AttachRoutes(mux, routes.Routes)
	config := config.LoadConfig()

	log.Fatal(http.ListenAndServe(config.Host + ":" + config.Port, mux))
}

