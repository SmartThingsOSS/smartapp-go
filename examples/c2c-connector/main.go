package main

import (
	"github.com/SmartThingsOSS/smartapp-go/examples/c2c-connector/internal/application"
	"log"
	"net/http"
)

var (
	version = "dev"
)

func getAddr(app application.Application) string {
	config := app.GetConfig()
	host := config.GetString("server.host")
	port := config.GetString("server.port")
	log.Printf("Starting on http://%s:%s", host, port)
	return host + ":" + port
}

func main() {
	app := application.GetInstance(version)
	engine := app.GetEngine()

	engine.Run(getAddr(app))

	server := &http.Server{
		Addr:         getAddr(app),
		Handler:      engine,
		ReadTimeout:  app.GetConfig().GetDuration("server.readTimeout"),
		WriteTimeout: app.GetConfig().GetDuration("server.writeTimeout"),
	}
	server.ListenAndServe()
}
