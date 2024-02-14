package routes

import (
	"fmt"
	"log"

	"spw-booking-service/src/api/v1/controller/booking"
	"spw-booking-service/src/utils"

	fiber "github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/monitor"
	"github.com/gofiber/fiber/v2/middleware/recover"
)

func StartServ() {
	fmt.Println("--- StartServ ---")
	iMiddleware := utils.NewMiddleware()

	app := fiber.New()
	app.Use(recover.New(
		recover.Config{
			EnableStackTrace: true,
		},
	))
	app.Use(cors.New())
	app.Use(iMiddleware.InitialMiddlewareRequestId("X-Request-ID"))
	systemConfig, err := utils.ReadConfig("configs")

	if err != nil {
		fmt.Println("read config fail")
	}

	app.Get("/dashboard", monitor.New())

	booking := booking.NewServer()
	booking.RegisterRoutes(app, &systemConfig)

	//Show Route
	for _, routes := range app.Stack() {
		for _, route := range routes {
			if route.Method == fiber.MethodGet || route.Method == fiber.MethodPost {
				fmt.Println(route.Method + ":" + route.Path)
			}
		}
	}

	log.Fatal(app.Listen(":" + systemConfig.Service.Port))

}
