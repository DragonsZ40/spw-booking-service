#=============================================================
#--------------------- build stage ---------------------------
#=============================================================
FROM golang:1.22.0-alpine AS build_stage
RUN apk update && apk add git
ENV PACKAGE_PATH=spw/ms/booking-service
RUN mkdir -p /go/src/
WORKDIR /go/src/$PACKAGE_PATH

ENV PATH="/usr/local/go/bin:$PATH"
COPY . /go/src/$PACKAGE_PATH/
RUN go clean -modcache
RUN go mod tidy
RUN go build -o ms-booking-service
#=============================================================
#--------------------- final stage ---------------------------
#=============================================================
FROM golang:1.22.0-alpine AS final_stage

ENV PACKAGE_PATH=spw/ms/booking-service

ENV http_proxy=
ENV https_proxy=

COPY --from=build_stage /go/src/$PACKAGE_PATH/ms-booking-service /go/src/$PACKAGE_PATH/
COPY --from=build_stage /go/src/$PACKAGE_PATH/configs /go/src/$PACKAGE_PATH/configs

WORKDIR /go/src/$PACKAGE_PATH/

ENTRYPOINT ./ms-booking-service local
EXPOSE 8080
