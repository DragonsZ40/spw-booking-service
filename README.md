**Go version 1.22**

1. You can use Dockerfile for run this project
2. docker build -t ms-booking-service .
3. docker run -p 8081:8080 -d ms-booking-service:latest

** API Test **
1. You can import `spwBookingService.postman_collection.json` into Postman for API Testing
2. config enviroment 
	base_url : http://localhost
	port : 8081