package boot

import (
	"github.com/kelseyhightower/envconfig"
	"log"
)

type AppConfig struct {
	AppName     string `envconfig:"go_application_name" required:"true"`
	RestConfig  *RestConfig
	MongoConfig *MongoConfig
}

type RestConfig struct {
	BaseUrl    string `envconfig:"rest_base_url" required:"true"`
	BasePort   string `envconfig:"rest_base_port" required:"true"`
	AuthHeader string `envconfig:"rest_auth_header" required:"true"`
}

type MongoConfig struct {
	Uri      string `envconfig:"mongodb_uri" required:"true"`
	Database string `envconfig:"mongodb_db" required:"true"`
}

func readConfig[T any]() *T {
	var config T
	if err := envconfig.Process("", &config); err != nil {
		log.Fatalf("Could not read config: %v", err)
	}
	return &config
}
