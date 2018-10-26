package config

import (
	log "github.com/sirupsen/logrus"
	"strings"
	"sync"

	"github.com/spf13/viper"
)

var instance *viper.Viper
var once sync.Once

func GetInstance() *viper.Viper {
	once.Do(func() {
		replacer := strings.NewReplacer(".", "_")
		v := viper.New()
		v.SetEnvPrefix("smartappgo")
		v.SetEnvKeyReplacer(replacer)
		v.SetConfigType("yaml")
		v.SetConfigName("smartapp-go")
		v.AddConfigPath("/etc/smartapp-go/")
		v.AddConfigPath("$HOME/.smartapp-go")
		v.AddConfigPath(".")
		v.AutomaticEnv()
		err := v.ReadInConfig()

		if err != nil {
			log.WithFields(log.Fields{
				"error": err,
			}).Fatal("configuration_error")
		}
		instance = v
	})
	return instance
}
