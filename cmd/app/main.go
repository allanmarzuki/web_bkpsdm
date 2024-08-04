package main

import (
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/template/html/v2"
	"github.com/allanmarzuki/web_bkpsdm.git/internal/auth"
	"github.com/allanmarzuki/web_bkpsdm.git/internal/database"
)

func main() {
	database.ConnectDB()

	app := fiber.New(fiber.Config{
		Views: html.New("./web/template", ".html"),
	})

	app.Static("/static", "./web/static")

	authRepo := auth.NewRepository()
	authService := auth.NewService(authRepo)
	authHandler := auth.NewHandler(authService)

	app.Get("/", func(c *fiber.Ctx) error {
		return c.Redirect("/login")
	})
	
	app.Get("/login", func(c *fiber.Ctx) error {
		return c.Render("login", fiber.Map{})
	})

	api := app.Group("/api")
	api.Post("/login", authHandler.Login)
	api.Post("/register", authHandler.Register)

	log.Fatal(app.Listen(":3000"))
}