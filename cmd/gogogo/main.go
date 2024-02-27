package main

import (
	"fmt"
	fibhelp "github.com/aiteung/athelper/fiber"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/compress"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/recover"
	"gogogo/internal/router/v1"
	"log"
	"os"
)

func main() {
	app := fiber.New(fibhelp.GenerateConfig())
	app.Use(recover.New(recover.Config{
		EnableStackTrace: true,
	}))
	app.Use(compress.New(compress.Config{
		Level: 5,
	}))

	app.Use(logger.New(logger.Config{
		Format: "${status} - ${method} ${path}\n",
	}))

	port := os.Getenv("PORT")
	if port == "" {
		port = "5000"
	}

	errv2 := v1.RouterV1(app)
	if errv2 != nil {
		log.Fatalf("%+v", errv2)
		return
	}

	log.Fatal(app.Listen(fmt.Sprintf("%s:%s", "0.0.0.0", port)))
}
