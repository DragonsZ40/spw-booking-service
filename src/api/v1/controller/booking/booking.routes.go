package booking

import (
	"spw-booking-service/src/utils"

	"github.com/gofiber/fiber/v2"
)

func (server *Server) RegisterRoutes(app *fiber.App, conf *utils.SystemConfig) {
	r := app.Group(server.Config.Service.Endpoint + "/booking")
	r.Post("/initialize-all-tables", server.InitializeAllTables)
	r.Post("/reserve-table", server.ReserveTable)
	r.Post("/cancel-reservation", server.CancelReservation)
}
