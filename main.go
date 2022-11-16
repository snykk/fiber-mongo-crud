package main

import (
	"fmt"
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/recover"
	config "github.com/snykk/fiber-mongo-crud/config"
	"github.com/snykk/fiber-mongo-crud/controller"
	"github.com/snykk/fiber-mongo-crud/repository"
	"github.com/snykk/fiber-mongo-crud/service"
)

func init() {
	// Setup Config Env
	if err := config.InitializeAppConfig(); err != nil {
		log.Fatal(err)
	}
}

func main() {
	// Setup Mongo
	configMongo := config.ConfigMongo{
		MONGO_HOST:                 config.AppConfig.MONGO_HOST,
		MONGO_PORT:                 config.AppConfig.MONGO_PORT,
		MONGO_USERNAME:             config.AppConfig.MONGO_USERNAME,
		MONGGO_PASSWORD:            config.AppConfig.MONGGO_PASSWORD,
		MONGO_DATABASE:             config.AppConfig.MONGO_DATABASE,
		MONGO_POOL_MIN:             config.AppConfig.MONGO_POOL_MIN,
		MONGO_POOL_MAX:             config.AppConfig.MONGO_POOL_MAX,
		MONGO_MAX_IDLE_TIME_SECOND: config.AppConfig.MONGO_MAX_IDLE_TIME_SECOND,
	}
	db, err := configMongo.InitializeMongo()
	if err != nil {
		panic(err)
	}

	// Setup Repository
	todoRepository := repository.NewTodoRepository(db)

	// Setup Service
	todoService := service.NewTodoService(&todoRepository)

	// Setup Controller
	todoController := controller.NewTodoController(&todoService)

	// Setup Fiber
	app := fiber.New()
	app.Use(recover.New())

	// Setup Routing
	todoController.Route(app)

	// Start App
	if err = app.Listen(fmt.Sprintf(":%d", config.AppConfig.PORT)); err != nil {
		panic(err)
	}
}
