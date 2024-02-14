package booking

import (
	"encoding/json"
	"spw-booking-service/src/api/v1/models"
	"spw-booking-service/src/utils"
	"time"

	"spw-booking-service/src/logger"

	"github.com/go-playground/validator"
	"github.com/gofiber/fiber/v2"
)

func (server *Server) InitializeAllTables(c *fiber.Ctx) error {
	start := time.Now()
	var req models.BookingRequest
	if err := c.BodyParser(&req); err != nil {
		return utils.ResponseFailed(c, fiber.StatusBadRequest, "failed", err)

	}
	response := utils.ResponseStandard{}
	v := validator.New()
	err := v.Struct(req)
	if err != nil {
		return utils.ResponseFailed(c, fiber.StatusBadRequest, "Required Field", err)
	}
	getLogModelRequest := logger.LogModelRequest{
		ServiceType:   "InitializeAllTables",
		CorrelationID: c.GetRespHeader("X-Request-Id"),
		Method:        "POST",
		StepName:      "InitializeAllTables",
		Start:         start,
		Suffix:        server.Config.Log.Kibana.Suffix,
	}
	logModel := logger.GetLogModel(getLogModelRequest)

	defer func() {
		logReq := req
		bt, _ := json.Marshal(logReq)
		logOrderRequest := logger.LogOrderRequest{
			Request:    string(bt),
			Response:   string(c.Response().Body()),
			ResultCode: response.Code,
			ResultDesc: response.Message,
		}
		logger.LogOrder(logOrderRequest, logModel, start)
	}()

	reqByte, _ := json.Marshal(req)
	loggerModel := logger.LogModel{
		ServiceType: "InitializeAllTables",
		Method:      "POST",
		StepName:    "InitializeAllTables",
		Txid:        c.GetRespHeader("X-Request-Id"),
		Request:     string(reqByte),
		Suffix:      server.Config.Log.Kibana.Suffix,
	}
	response = server.api.ServiceBooking.InitializeAllTablesService(req, loggerModel)
	if response.Code == "200" {
		return utils.ResponseSuccess(c, response.Message, response.Result)
	} else {
		return utils.ResponseFailed(c, fiber.StatusBadRequest, response.Message, err)
	}
}

func (server *Server) ReserveTable(c *fiber.Ctx) error {
	start := time.Now()
	var req models.ReserveTableRequest
	if err := c.BodyParser(&req); err != nil {
		return utils.ResponseFailed(c, fiber.StatusBadRequest, "failed", err)

	}
	response := utils.ResponseStandard{}
	v := validator.New()
	err := v.Struct(req)
	if err != nil {
		return utils.ResponseFailed(c, fiber.StatusBadRequest, "Required Field", err)
	}
	getLogModelRequest := logger.LogModelRequest{
		ServiceType:   "ReserveTable",
		CorrelationID: c.GetRespHeader("X-Request-Id"),
		Method:        "POST",
		StepName:      "ReserveTable",
		Start:         start,
		Suffix:        server.Config.Log.Kibana.Suffix,
	}
	logModel := logger.GetLogModel(getLogModelRequest)

	defer func() {
		logReq := req
		bt, _ := json.Marshal(logReq)
		logOrderRequest := logger.LogOrderRequest{
			Request:    string(bt),
			Response:   string(c.Response().Body()),
			ResultCode: response.Code,
			ResultDesc: response.Message,
		}
		logger.LogOrder(logOrderRequest, logModel, start)
	}()

	reqByte, _ := json.Marshal(req)
	loggerModel := logger.LogModel{
		ServiceType: "ReserveTable",
		Method:      "POST",
		StepName:    "ReserveTable",
		Txid:        c.GetRespHeader("X-Request-Id"),
		Request:     string(reqByte),
		Suffix:      server.Config.Log.Kibana.Suffix,
	}
	response = server.api.ServiceBooking.ReserveTableService(req, loggerModel)
	if response.Code == "200" {
		return utils.ResponseSuccess(c, response.Message, response.Result)
	} else {
		return utils.ResponseFailed(c, fiber.StatusBadRequest, response.Message, nil)
	}
}

func (server *Server) CancelReservation(c *fiber.Ctx) error {
	start := time.Now()
	var req models.CancelBookingRequest
	if err := c.BodyParser(&req); err != nil {
		return utils.ResponseFailed(c, fiber.StatusBadRequest, "failed", err)

	}
	response := utils.ResponseStandard{}
	v := validator.New()
	err := v.Struct(req)
	if err != nil {
		return utils.ResponseFailed(c, fiber.StatusBadRequest, "Required Field", err)
	}
	getLogModelRequest := logger.LogModelRequest{
		ServiceType:   "CancelReservation",
		CorrelationID: c.GetRespHeader("X-Request-Id"),
		Method:        "POST",
		StepName:      "CancelReservation",
		Start:         start,
		Suffix:        server.Config.Log.Kibana.Suffix,
	}
	logModel := logger.GetLogModel(getLogModelRequest)

	defer func() {
		logReq := req
		bt, _ := json.Marshal(logReq)
		logOrderRequest := logger.LogOrderRequest{
			Request:    string(bt),
			Response:   string(c.Response().Body()),
			ResultCode: response.Code,
			ResultDesc: response.Message,
		}
		logger.LogOrder(logOrderRequest, logModel, start)
	}()

	reqByte, _ := json.Marshal(req)
	loggerModel := logger.LogModel{
		ServiceType: "CancelReservation",
		Method:      "POST",
		StepName:    "CancelReservation",
		Txid:        c.GetRespHeader("X-Request-Id"),
		Request:     string(reqByte),
		Suffix:      server.Config.Log.Kibana.Suffix,
	}
	response = server.api.ServiceBooking.CancelReservationTableService(req, loggerModel)
	if response.Code == "200" {
		return utils.ResponseSuccess(c, response.Message, response.Result)
	} else {
		return utils.ResponseFailed(c, fiber.StatusBadRequest, response.Message, nil)
	}
}
