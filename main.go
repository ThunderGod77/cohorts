package main

import (
	"cohort/db"
	"cohort/global"
	"cohort/routes"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"log"
	"os"
	"os/signal"
)

func setupRoutes(app *fiber.App) {
	v1 := app.Group("/v3")
	routes.StudentRoutes(v1.Group("/student"))
}

func main() {

	//to initialize postgres database
	db.Init()
	defer global.Dbpool.Close()

	//to initialize warning,info and error loggers
	global.Init()
	// to close logs.txt file
	defer global.File.Close()

	app := fiber.New(fiber.Config{
		// Override default error handler
		ErrorHandler: global.ErrHandler,
	})

	// to enable cors
	app.Use(cors.New(cors.Config{
		AllowOrigins: "*",
		AllowHeaders: "Origin, Content-Type, Accept, Authorization",
	}))

	//to enable for logging to standard console
	app.Use(logger.New())
	//to enable logging to logs.txt file
	app.Use(logger.New(logger.Config{
		Output: global.File,
	}))

	//test route
	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Hello, World!")
	})

	//sets up routes to the rest api
	setupRoutes(app)

	//to enable graceful shutdown
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)
	go func() {
		_ = <-c
		global.WarningLogger.Println("Gracefully shutting down...")
		log.Println("Gracefully shutting down...")
		_ = app.Shutdown()
	}()

	//starts the server to listen to routes
	err := app.Listen(":8080")
	if err != nil {
		log.Println(err)
	}
}
