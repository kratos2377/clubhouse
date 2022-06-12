package main

import (
	"flag"
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/kratos2377/clubhouse/middleware"
	"github.com/kratos2377/clubhouse/routes"
)

func main() {
	app := fiber.New()

	middleware.SetMiddleware(app)
	routes.SetupApiV1(app)

	port := "8000"

	addr := flag.String("addr", port, "http service address")
	log.Fatal(app.Listen(":" + *addr))
}
