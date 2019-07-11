package application

import (
	"encoding/json"
	"github.com/SmartThingsOSS/smartapp-go/examples/c2c-connector/internal/config"
	"github.com/SmartThingsOSS/smartapp-go/examples/c2c-connector/internal/handlers"
	"github.com/SmartThingsOSS/smartapp-go/examples/c2c-connector/internal/logging"
	"github.com/SmartThingsOSS/smartapp-go/pkg/smartappcore"
	"github.com/SmartThingsOSS/smartapp-go/pkg/smartappgin"
	log "github.com/sirupsen/logrus"
	"sync"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
)

type Application interface {
	GetEngine() *gin.Engine
	GetConfig() *viper.Viper
}

type SmartApp struct {
	logger *log.FieldLogger
	engine *gin.Engine
	config *viper.Viper
}

var instance Application
var once sync.Once

func GetInstance(version string) Application {
	once.Do(func() {
		log.SetFormatter(&log.TextFormatter{
			DisableTimestamp:       false,
			FullTimestamp:          true,
			DisableLevelTruncation: true,
			DisableColors:          true,
		})
		instance = application()
		appConfig := instance.GetConfig()
		log.SetLevel(getLogLevel(appConfig.GetString("server.logLevel")))
		engine := instance.GetEngine()
		engine.Use(logging.LoggingContext())
		engine.Use(logging.RequestLogger(time.RFC3339, true))

		engine.GET("/health", func(c *gin.Context) {
			c.JSON(200, gin.H{})
		})

		engine.GET("/buildinfo", func(c *gin.Context) {
			c.JSON(200, gin.H{
				"version": version,
			})
		})

		engine.POST("/", smartapp(appConfig))
	})

	return instance
}

func application() Application {
	engine := gin.New()
	appConfig := config.GetInstance()

	return &SmartApp{
		engine: engine,
		config: appConfig,
	}
}

func (s *SmartApp) GetEngine() *gin.Engine {
	return s.engine
}

func (s *SmartApp) GetConfig() *viper.Viper {
	return s.config
}

func getLogLevel(val string) log.Level {
	level, err := log.ParseLevel(val)
	if err != nil {
		log.WithFields(log.Fields{
			"configuredLevel": val,
		}).Panic("Invalid configured log level. ", err)
	}
	return level
}

func smartapp(config *viper.Viper) gin.HandlerFunc {
	var authConfig smartappcore.AuthConfig
	if err := config.UnmarshalKey("auth", &authConfig); err != nil {
		log.WithFields(log.Fields{
			"error": err,
		}).Panic("Unable to parse auth config. ", err)
	}
	api := smartappcore.NewSmartThingsApi(&smartappcore.SmartThingsApiParams{
		Host:    config.GetString("smartthings.host"),
		Schemes: []string{config.GetString("smartthings.scheme")},
		Debug:   config.GetBool("smartthings.debug"),
		Logger:  log.StandardLogger(),
	})

	definition := smartappcore.NewSmartAppDefinition()
	definition.SetInstallHandler(handlers.NewInstallHandler(api).Handler())
	definition.SetEventHandler(handlers.NewEventHandler(api).Handler())
	definition.SetConfigurationHandler(handlers.NewConfigHandler().Handler())
	definition.SetAuthConfig(&authConfig)
	app := smartappgin.NewSmartApp(definition)
	app.SetRequestInterceptor(getRequestInterceptor())

	return app.Handler()
}

func getRequestInterceptor() smartappgin.RequestInterceptor {
	return func(params *smartappcore.SmartAppParams) {
		request := params.Request
		ctx := params.Context
		logger := logging.Logger(ctx)
		logger.WithFields(log.Fields{
			"lifecycle": request.Lifecycle,
			"request":   prettyPrint(request),
		}).Debug("smartapp_request_received")
	}
}

func prettyPrint(i interface{}) string {
	s, _ := json.Marshal(i)
	return string(s)
}
