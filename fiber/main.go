package main

import (
	"fmt"
	"github.com/gofiber/fiber"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
	"github.com/mahesh-dilhan/go-gRPC-microservices/country"
	"github.com/mahesh-dilhan/go-gRPC-microservices/database"
)

func initDatabase() {
	var err error
	database.DBConn, err = gorm.Open("sqlite3", "countries.db")
	if err != nil {
		panic("failed to connect database")
	}
	fmt.Println("Connection Opened to Database")
	database.DBConn.AutoMigrate(&country.Country{})

}

func helloFromFiber(c *fiber.Ctx) {
	c.Send("Hello, World!")
}

func routers(app *fiber.App) {
	app.Get("/", helloFromFiber)
	app.Get("/api/v1/countries", country.GetCountries)
	app.Get("/api/v1/country/:id", country.GetCountry)
	app.Post("/api/v1/country", country.NewCountry)

}

func main() {

	app := fiber.New()
	initDatabase()
	routers(app)
	app.Listen(4003)
	defer database.DBConn.Close()
}
