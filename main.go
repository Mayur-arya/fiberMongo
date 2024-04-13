// Om Namah Shivay
package main

import (
	"fiberMongo/connection"
	"fiberMongo/routes"
	"fmt"

	"github.com/gofiber/fiber/v2"
)

func main() {
	// Create a Fiber app
	app := fiber.New()
	app.Post("/insertUser", routes.AddUser)
	connection.DbConnection()
	fmt.Println("This is for fiber api with mongo db..")
	app.Listen(":3000")
}
