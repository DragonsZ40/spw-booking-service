package booking

import (
	bookingServices "spw-booking-service/src/api/v1/services/booking"
	"spw-booking-service/src/utils"
)

type API struct {
	//ServiceBooking productService.IProductService
	ServiceBooking bookingServices.IBookingService
}

type Server struct {
	api    API
	Config utils.SystemConfigDatabaseList
}

func NewServer() *Server {

	confSystem, err := utils.ReadConfigDatabaseList("configs")
	if err != nil {
		panic(err)
	}

	// iProduct := productService.NewIServiceProduct(productService.NewIServiceDaoProduct(systemDB))

	// return &Server{
	// 	Config: confSystem,
	// 	api: API{
	// 		ServiceProduct: iProduct,
	// 	},
	// }

	iBooking := bookingServices.NewIServiceProduct()

	return &Server{
		Config: confSystem,
		api: API{
			ServiceBooking: iBooking,
		},
	}
}
