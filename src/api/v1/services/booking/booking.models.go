package bookingServices

type Table struct {
	ID               int  `json:"id"`
	IsReserved       bool `json:"isReserved"`
	ReservedForGroup int  `json:"reservedForGroup"`
}
