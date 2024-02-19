package main

import (
	"fmt"
	"log"
	"os"
	"path/filepath"

	"github.com/beego/beego/v2/client/orm"

	"github.com/Yasoobnaqvi/v4tech_task/v4_tech.git/database"
	"github.com/Yasoobnaqvi/v4tech_task/v4_tech.git/models"
	"github.com/Yasoobnaqvi/v4tech_task/v4_tech.git/routes"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/joho/godotenv"
)

func main() {
	app := fiber.New()
	app.Use(cors.New(cors.Config{
		AllowOrigins: "*",
	}))

	SetupEnv()
	SetupRoutes(app)
	orm.RegisterModel(new(models.Users))
	database.ConnectDB()

	log.Fatal(app.Listen(fmt.Sprintf(":%s", os.Getenv("PORT"))))

}

func SetupRoutes(app fiber.Router) {
	app.Post("/products", routes.Signup)
	app.Get("/products", routes.Signup)
	app.Get("/products/:id/", routes.Signup)
	app.Put("/products/:id/", routes.Signup)
	app.Delete("/products/:id/", routes.Signup)
}

func SetupEnv() {
	if os.Getenv("environment") != "test" {
		if dir, err := filepath.Abs(filepath.Dir(os.Args[0])); err != nil {
			log.Fatal(err)
		} else {
			_ = godotenv.Load(dir + "/.env")
		}
	} else {
		if err := godotenv.Load(); err != nil {
			log.Fatal(err)
		}
	}
}