package country

import (
	"fmt"
	"github.com/gofiber/fiber"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
	"github.com/mahesh-dilhan/go-gRPC-microservices/database"
)

type Country struct {
	gorm.Model
	Country          string `json:"country"`
	NumberOfPatients int    `json:"numberOfPatients"`
}

func GetCountries(c *fiber.Ctx) {
	db := database.DBConn
	var countries []Country
	db.Find(&countries)
	c.JSON(countries)
}

func GetCountry(c *fiber.Ctx) {
	id := c.Params("id")
	db := database.DBConn
	var country Country
	db.Find(&country, id)
	c.JSON(country)
}

func NewCountry(c *fiber.Ctx) {
	db := database.DBConn
	country := new(Country)

	if err := c.BodyParser(country); err != nil {
		fmt.Println(err)
		panic("Unable to parse")
	}
	db.Create(&country)
	c.JSON(country)
}
