package models

type BookingRequest struct {
	NumberOfTables int `json:"numberOfTables" validate:"required"`
}

type ReserveTableRequest struct {
	NumberOfCustomer int `json:"numberOfCustomer" validate:"required"`
}

type CancelBookingRequest struct {
	BookingId int `json:"bookingId" validate:"required"`
}
