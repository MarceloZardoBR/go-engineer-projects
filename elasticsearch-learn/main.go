package main

import (
	"log"

	"github.com/MarceloZardoBR/go-api-frame/database"
	"github.com/MarceloZardoBR/go-api-frame/domain/services"
	externalservices "github.com/MarceloZardoBR/go-api-frame/external-services"
	"github.com/MarceloZardoBR/go-api-frame/infra/config"
	"github.com/MarceloZardoBR/go-api-frame/router"
	"github.com/MarceloZardoBR/go-api-frame/router/mainrouter/authorization"
	"github.com/MarceloZardoBR/go-api-frame/router/mainrouter/titlessearchengine"
	"github.com/gofiber/fiber/v2"

	"github.com/MarceloZardoBR/go-api-frame/server"
)

func main() {

	config, err := config.ReadAndLoadEnvVars()
	if err != nil {
		log.Println(err)
	}

	db, err := database.Instance(config)
	if err != nil {
		log.Println(err)
	}

	_ = services.NewServices(db, config)

	elasticSearchService := externalservices.NewElasticSearchService(config)

	titlesSearchController := titlessearchengine.NewController(elasticSearchService)
	authController := authorization.NewController(config)

	fiber := fiber.New()
	server := server.NewServer(fiber)

	router.AddRouter(fiber, authController, titlesSearchController)

	server.StartServer("5000")
}
