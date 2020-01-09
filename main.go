package main

import (
	kitHttp "github.com/go-kit/kit/transport/http"
	"github.com/labstack/echo/v4"
	echoSwagger "github.com/swaggo/echo-swagger"

	"go-pb-swagger/greeting"
)

var svc greeting.GreetingService

// @title Hello API
// @version 0.1
// @description

// @contact.name n
// @contact.email navono007@gmail.com

// @host 10.10.10.10:10888
// @BasePath /api
func main() {
	svc = greeting.NewGreetingService()

	e := echo.New()
	e.GET("/", echo.WrapHandler(kitHttp.NewServer(
		greeting.MakeGreetingEp(svc),
		greeting.DecHelloRequest,
		greeting.EncHelloRequest,
	)))

	e.GET("/swagger", echoSwagger.WrapHandler)

	e.Logger.Fatal(e.Start(":10086"))
}
