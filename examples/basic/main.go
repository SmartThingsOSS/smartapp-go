package main

import (
	"encoding/json"
	"errors"
	"github.com/SmartThingsOSS/smartapp-go/pkg/smartapp"
	"github.com/SmartThingsOSS/smartapp-go/pkg/smartappcore"
	"github.com/SmartThingsOSS/smartapp-go/pkg/smartappgin"
	"github.com/gin-gonic/gin"
	"log"
)

func main() {
	r := setupRouter()
	log.Printf("Starting on http://localhost:5555")
	r.Run(":5555")
}

func setupRouter() *gin.Engine {
	r := gin.Default()

	r.POST("/", initSmartapp())

	return r
}

func initSmartapp() gin.HandlerFunc {
	definition := smartappcore.NewSmartAppDefinition()
	definition.SetInstallHandler(func(params *smartappcore.SmartAppParams) (*smartapp.ExecutionResponse, error) {
		log.Printf("Handling Install lifecycle")
		return &smartapp.ExecutionResponse{StatusCode: 200}, nil
	})
	definition.SetEventHandler(func(params *smartappcore.SmartAppParams) (*smartapp.ExecutionResponse, error) {
		log.Printf("Handling Event lifecycle")
		return &smartapp.ExecutionResponse{StatusCode: 200}, nil
	})
	definition.SetConfigurationHandler(configHandler)

	definition.SetAuthConfig(&smartappcore.AuthConfig{
		PublicKeyBase64: "",
	})
	app := smartappgin.NewSmartApp(definition)

	app.SetRequestInterceptor(getRequestInterceptor())

	return app.Handler()
}

func configHandler(params *smartappcore.SmartAppParams) (*smartapp.ExecutionResponse, error) {
	log.Printf("Handling Config lifecycle")
	switch params.Request.ConfigurationData.Phase {
	case smartapp.ConfigurationPhaseINITIALIZE:
		return &smartapp.ExecutionResponse{
			StatusCode: 200,
			ConfigurationData: &smartapp.ConfigurationResponseData{
				Initialize: &smartapp.InitializeSetting{
					Setting: smartapp.Setting{
						Name:        "Basic SmartApp",
						Description: "Basic SmartApp",
						ID:          "app",
					},
					FirstPageID: "1",
					Permissions: []string{"r:devices:*", "w:devices:*", "x:devices:*"},
				},
			},
		}, nil
	case smartapp.ConfigurationPhasePAGE:
		return &smartapp.ExecutionResponse{
			StatusCode: 200,
			ConfigurationData: &smartapp.ConfigurationResponseData{
				Page: &smartapp.Page{
					PageID:   "1",
					Name:     "Basic SmartApp",
					Complete: boolPointer(true),
					Sections: []*smartapp.Section{
						{
							Settings: []*smartapp.SectionSetting{
								{
									Setting: smartapp.Setting{
										Name:        "Basic SmartApp",
										Description: "Pretty basic huh?!  I wasn't kidding.",
										ID:          "myParagraphSetting",
									},
									Type: smartapp.SettingTypePARAGRAPH,
								},
							},
						},
					},
				},
			},
		}, nil
	default:
		return &smartapp.ExecutionResponse{}, errors.New("UnsupportedConfigurationPhaseError")
	}
}

func boolPointer(val bool) *bool {
	return &val
}

func getRequestInterceptor() smartappgin.RequestInterceptor {
	return func(params *smartappcore.SmartAppParams) {
		log.Printf("incoming_request lifecycle=%v body=%v", params.Request.Lifecycle, prettyPrint(params.Request))
	}
}

func prettyPrint(i interface{}) string {
	s, _ := json.Marshal(i)
	return string(s)
}
