package main

import (
	"the-antichrist-imperium/goat"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/template/django/v3"
)

func IndexView(c *fiber.Ctx) error {
	return c.Render("index", fiber.Map{}, "_layouts/_base")
}

func MarkDownConvert(c *fiber.Ctx) error {
	source := c.FormValue("markdown_source")
	response := goat.ExportHtml(source)
	return c.Render("markdown_result", fiber.Map{"response": response})
}

func main() {
	engine := django.New("./views", ".html")
	engine.Debug(true)

	app := fiber.New(fiber.Config{
		Views: engine,
	})

	app.Static("/static/", "./static")

	app.Get("/", IndexView)
	app.Post("/mark/", MarkDownConvert)

	app.Listen(":3000")
}
