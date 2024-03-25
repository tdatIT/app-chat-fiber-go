package config

import (
	"errors"
	"fmt"
	"github.com/google/wire"
	"github.com/spf13/viper"
	"log"
	"os"
)

var Set = wire.NewSet(NewConfig)

type AppConfig struct {
	DebugMode bool
	BaseURL   string
	SecretKey string
	Database  Database
	RestAPI   RestAPI
	WebSocket Websocket
	Version   string
}

type Database struct {
	MongoDB MongoDB
}

type MongoDB struct {
	URI               string
	Database          string
	ConnectionTimeout int
	Timeout           int
	MaxConnIdleTime   int
}

type RestAPI struct {
	Port                string
	Mode                string
	SSL                 bool
	CtxDefaultTimeout   int
	ExpirationLimitTime int
	ReadTimeout         int
	WriteTimeout        int
}

type Websocket struct {
	Port            string
	ReadBufferSize  int
	WriteBufferSize int
	WriteWait       int
	PongWait        int
	PingPeriod      int
	MaxMessageSize  int
}

func getDefaultConfig() string {
	return "/config/config"
}
func NewConfig() (*AppConfig, error) {
	cfg := AppConfig{}
	path := os.Getenv("cfgPath")
	if path == "" {
		path = getDefaultConfig()
	}
	fmt.Printf("config path:%s\n", path)

	v := viper.New()
	v.SetConfigFile(path)
	v.AutomaticEnv()

	if err := v.ReadInConfig(); err != nil {
		var configFileNotFoundError viper.ConfigFileNotFoundError
		if errors.As(err, &configFileNotFoundError) {
			return nil, errors.New("config file not found")
		}
		return nil, err
	}

	err := v.Unmarshal(&cfg)
	if err != nil {
		log.Printf("unable to decode into struct, %v", err)
		return nil, err
	}

	return &cfg, nil
}
