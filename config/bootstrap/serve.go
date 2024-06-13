package bootstrap

import (
	"go_first_project/pkg/config"
	"go_first_project/pkg/html"
	"go_first_project/pkg/routing"
)

func Serve() {
	config.Set()

	routing.Init()

	html.LoadHTML(routing.GetRouter())

	routing.RegisterRoutes()

	routing.Serve()

}
