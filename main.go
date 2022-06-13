package main

import (
	"context"
	"flag"
	"fmt"
	"log"

	"entgo.io/ent/entc/integration/customid/ent"
	"entgo.io/ent/entc/integration/ent/migrate"
	"github.com/gofiber/fiber/v2"
	"github.com/kratos2377/clubhouse/config"
	"github.com/kratos2377/clubhouse/middleware"
	"github.com/kratos2377/clubhouse/routes"
	"github.com/kratos2377/clubhouse/utils"
)

func main() {
	conf := config.New()
	client, err := ent.Open("postgres", fmt.Sprintf("host=%s port=%s user=%s dbname=%s password=%s sslmode=disable", conf.Database.Host, conf.Database.Port, conf.Database.User, conf.Database.Name, conf.Database.Password))
	if err != nil {
		utils.Fatalf("Database connection failed: ", err)
	}
	defer client.Close()

	ctx := context.Background()

	err = client.Schema.Create(
		ctx,
		migrate.WithDropIndex(true),
		migrate.WithDropColumn(true),
	)

	if err != nil {
		utils.Fatalf("Migration fail: ", err)
	}

	app := fiber.New()

	middleware.SetMiddleware(app)
	routes.SetupApiV1(app)

	port := "8000"

	addr := flag.String("addr", port, "http service address")
	log.Fatal(app.Listen(":" + *addr))
}
