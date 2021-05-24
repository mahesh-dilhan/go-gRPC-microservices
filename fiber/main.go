package main

import "github.com/gofiber/fiber"

func helloFromFiber(c *fiber.Ctx) {
	c.Send("Hello, World!")
}

func routers(app *fiber.App) {
	app.Get("/", helloFromFiber)
}

func main() {

	app := fiber.New()
	routers(app)
	app.Listen(4003)
}
