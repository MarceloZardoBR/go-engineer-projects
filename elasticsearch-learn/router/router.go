package router

import (
	"github.com/MarceloZardoBR/go-api-frame/router/mainrouter/authorization"
	"github.com/MarceloZardoBR/go-api-frame/router/mainrouter/titlessearchengine"
	"github.com/MarceloZardoBR/go-api-frame/server/middlewares"
	"github.com/gofiber/fiber/v2"
)

func AddRouter(app *fiber.App, authController *authorization.Controller, titlesSearchController *titlessearchengine.Controller) {

	rootGroup := app.Group("")

	rootGroup.Get("/health/", func(ctx *fiber.Ctx) error {
		return ctx.JSON("OK")
	})

	mainGroup := rootGroup.Group("main")
	mainGroup.Get("/authorization/", authController.Authorization)

	searchGroup := rootGroup.Group("search")
	searchGroup.Get("/:origen/titles", titlesSearchController.SearchIndexWithQuery)

	authGroup := mainGroup.Group("")

	authGroup.Use(middlewares.Authorization())
}
