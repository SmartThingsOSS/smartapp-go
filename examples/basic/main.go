package main

import (
	"github.com/SmartThingsOSS/smartapp-go/pkg/smartapp"
	"github.com/SmartThingsOSS/smartapp-go/pkg/smartappcore"
	"github.com/SmartThingsOSS/smartapp-go/pkg/smartappgin"
	"github.com/gin-gonic/gin"
	"log"
)

func setupRouter() *gin.Engine {
	r := gin.Default()

	r.POST("/", initSmartapp())

	return r
}

func initSmartapp() gin.HandlerFunc {
	publicKey := ""
	definition := smartappcore.NewSmartAppDefinition()
	definition.SetInstallHandler(func(params *smartappcore.SmartAppParams) (*smartapp.ExecutionResponse, error) {
		log.Printf("Handling Install lifecycle")
		return &smartapp.ExecutionResponse{StatusCode: 200}, nil
	})
	definition.SetEventHandler(func(params *smartappcore.SmartAppParams) (*smartapp.ExecutionResponse, error) {
		log.Printf("Handling Event lifecycle")
		return &smartapp.ExecutionResponse{StatusCode: 200}, nil
	})
	definition.SetConfigurationHandler(func(params *smartappcore.SmartAppParams) (*smartapp.ExecutionResponse, error) {
		log.Printf("Handling Config lifecycle")
		return &smartapp.ExecutionResponse{StatusCode: 200}, nil
	})
	app := smartappgin.NewSmartApp(definition)

	if publicKey != "" {
		app.SetPublicKey(smartappcore.ParsePublicKeyBase64(publicKey))
	}

	return app.Handler()
}

func main() {
	r := setupRouter()
	log.Printf("Starting on http://localhost:5555")
	r.Run(":5555")
}
