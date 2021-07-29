package main

import (
	"context"
	"log"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/wonpanu/my-golang/pkg/repo"
	handler "github.com/wonpanu/my-golang/pkg/route"
	"github.com/wonpanu/my-golang/pkg/usecase"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

func connectMongoDB() (*mongo.Client, error) {
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)

	client, err := mongo.Connect(ctx, options.Client().ApplyURI("mongodb://root:example@localhost:27017"))
	if err != nil {
		log.Fatal(err)
	}

	err = client.Ping(ctx, readpref.Primary())
	if err != nil {
		log.Fatal(err)
	}

	log.Println("Database connected 🎉")
	return client, err
}

func setupRoutes(app *fiber.App, mongo *mongo.Client) {

	// repo->usecase->route/adapter
	vaccineRepo := repo.NewVaccineRepo(mongo)
	vaccineUC := usecase.NewVaccineUsecase(vaccineRepo)
	vacecineHandler := handler.NewVaccineHandler(vaccineUC)

	app.Get("/", func(c *fiber.Ctx) error {
		return c.Status(fiber.StatusOK).JSON(fiber.Map{
			"success": true,
			"message": "You are at the endpoint 😉",
		})
	})

	api := app.Group("/api/v1")
	api.Get("/vaccine", vacecineHandler.GetAllVaccine)
	api.Get("/vaccine/:id", vacecineHandler.GetVaccineByID)
	api.Post("/vaccine", vacecineHandler.CreateVaccine)
	api.Put("/vaccine/:id", vacecineHandler.UpdateVaccine)
	api.Delete("/vaccine/:id", vacecineHandler.DeleteVaccine)
}

func main() {
	app := fiber.New()
	app.Use(logger.New())

	client, err := connectMongoDB()
	defer client.Disconnect(context.TODO())
	if err != nil {
		log.Println("Failed to disconnect database!")
	}

	setupRoutes(app, client)

	err = app.Listen(":8000")
	if err != nil {
		log.Fatal("Failed to start")
		panic(err)
	}
}
