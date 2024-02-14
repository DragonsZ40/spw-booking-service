package bookingServices

import (
	"encoding/json"
	"spw-booking-service/src/api/v1/models"
	"spw-booking-service/src/logger"
	"spw-booking-service/src/utils"
	"strconv"
	"sync"
	"time"
)

type IBookingService interface {
	InitializeAllTablesService(models.BookingRequest, logger.LogModel) (response utils.ResponseStandard)
	ReserveTableService(models.ReserveTableRequest, logger.LogModel) (response utils.ResponseStandard)
	CancelReservationTableService(models.CancelBookingRequest, logger.LogModel) (response utils.ResponseStandard)
}

type serviceBooking struct{}

func NewIServiceProduct() IBookingService {
	return &serviceBooking{}
}

var (
	tables           []Table
	isInitialized    bool
	bookingIDCounter int
	mutex            sync.Mutex // For concurrency control
)

func (r *serviceBooking) InitializeAllTablesService(req models.BookingRequest, log logger.LogModel) (response utils.ResponseStandard) {
	mutex.Lock()
	defer mutex.Unlock()

	startStep := time.Now()

	reqByte, _ := json.Marshal(req)
	logStepRequest := logger.LogStepRequest{
		StepName:    "InitializeAllTablesService",
		StartDate:   utils.ConvDatetimeFormatLog(startStep),
		StepRequest: string(reqByte),
		Endpoint:    "SPW-BOOKING-SERVIC",
		Method:      "InitializeAllTablesService",
		System:      "SPW-BOOKING-SERVIC",
	}

	if isInitialized {
		logStepRequest.ResultDesc = "Tables have already been initialized."
		logStepRequest.ResultCode = "400"
		logger.LogStep(logStepRequest, log, startStep)

		response.Code = "400"
		response.Result = nil
		response.Message = "Tables have already been initialized."
		return
	}

	tables = make([]Table, req.NumberOfTables)
	for i := range tables {
		tables[i] = Table{ID: i + 1, IsReserved: false}
	}
	isInitialized = true
	bookingIDCounter = 1

	response.Code = "200"
	response.Result = nil
	response.Message = "Tables initialized successfully."
	return

}

func (r *serviceBooking) ReserveTableService(req models.ReserveTableRequest, log logger.LogModel) (response utils.ResponseStandard) {
	mutex.Lock()
	defer mutex.Unlock()

	startStep := time.Now()

	reqByte, _ := json.Marshal(req)
	logStepRequest := logger.LogStepRequest{
		StepName:    "ReserveTableService",
		StartDate:   utils.ConvDatetimeFormatLog(startStep),
		StepRequest: string(reqByte),
		Endpoint:    "SPW-BOOKING-SERVIC",
		Method:      "ReserveTableService",
		System:      "SPW-BOOKING-SERVIC",
	}

	if !isInitialized {
		logStepRequest.ResultDesc = "Tables is not initialized. Please initialize tables first."
		logStepRequest.ResultCode = "400"
		logger.LogStep(logStepRequest, log, startStep)

		response.Code = "400"
		response.Result = nil
		response.Message = "Tables is not initialized. Please initialize tables first."
		return
	}

	tablesNeeded := (req.NumberOfCustomer + 3) / 4 // Ceiling division for groups > 4
	reservedTables := make([]int, 0)

	for i := 0; i < len(tables) && len(reservedTables) < tablesNeeded; i++ {
		if !tables[i].IsReserved {
			tables[i].IsReserved = true
			tables[i].ReservedForGroup = bookingIDCounter
			reservedTables = append(reservedTables, tables[i].ID)
		}
	}
	if len(reservedTables) < tablesNeeded {

		for i := range reservedTables {
			for j := range tables {
				if tables[j].ID == reservedTables[i] {
					tables[j].IsReserved = false
					tables[j].ReservedForGroup = 0
				}
			}
		}

		logStepRequest.ResultDesc = "Not enough tables available."
		logStepRequest.ResultCode = "400"
		logger.LogStep(logStepRequest, log, startStep)

		response.Code = "400"
		response.Result = nil
		response.Message = "Not enough tables available."
		return
	}

	remainingTables := len(tables) - len(reservedTables)

	resp := struct {
		BookingID       int   `json:"bookingId"`
		BookedTables    []int `json:"bookedTables"`
		RemainingTables int   `json:"remainingTables"`
	}{
		BookingID:       bookingIDCounter,
		BookedTables:    reservedTables,
		RemainingTables: remainingTables,
	}

	bookingIDCounter++

	response.Code = "200"
	response.Result = resp
	response.Message = "Success"
	return
}

func (r *serviceBooking) CancelReservationTableService(req models.CancelBookingRequest, log logger.LogModel) (response utils.ResponseStandard) {
	mutex.Lock()
	defer mutex.Unlock()

	startStep := time.Now()

	reqByte, _ := json.Marshal(req)
	logStepRequest := logger.LogStepRequest{
		StepName:    "CancelReservationTableService",
		StartDate:   utils.ConvDatetimeFormatLog(startStep),
		StepRequest: string(reqByte),
		Endpoint:    "SPW-BOOKING-SERVIC",
		Method:      "CancelReservationTableService",
		System:      "SPW-BOOKING-SERVIC",
	}

	if !isInitialized {
		logStepRequest.ResultDesc = "Tables is not initialized. Please initialize tables first."
		logStepRequest.ResultCode = "400"
		logger.LogStep(logStepRequest, log, startStep)

		response.Code = "400"
		response.Result = nil
		response.Message = "Tables is not initialized. Please initialize tables first."
		return
	}

	freedTablesCount := 0
	for i := range tables {
		if tables[i].ReservedForGroup == req.BookingId {
			tables[i].IsReserved = false
			tables[i].ReservedForGroup = 0
			freedTablesCount++
		}
	}

	if freedTablesCount == 0 {
		logStepRequest.ResultDesc = "Booking ID " + strconv.Itoa(req.BookingId) + " not found."
		logStepRequest.ResultCode = "400"
		logger.LogStep(logStepRequest, log, startStep)

		response.Code = "400"
		response.Result = nil
		response.Message = "Booking ID " + strconv.Itoa(req.BookingId) + " not found."
		return
	}

	remainingTables := len(tables) - countReservedTables()

	resp := struct {
		FreedTables     int `json:"freedTables"`
		RemainingTables int `json:"remainingTables"`
	}{
		FreedTables:     freedTablesCount,
		RemainingTables: remainingTables,
	}

	response.Code = "200"
	response.Result = resp
	response.Message = "Reservation cancelled successfully."
	return
}

func countReservedTables() int {
	count := 0
	for _, table := range tables {
		if table.IsReserved {
			count++
		}
	}
	return count
}
