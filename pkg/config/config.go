package config

import (
	"github.com/spf13/viper"
	"mango_crm/pkg/logger"
)

// Config agrupa la configuración de la aplicación.
// Cada etiqueta `mapstructure:"<KEY>"` indica a Viper (que usa
// mitchellh/mapstructure) qué nombre de clave mapear desde el config file o env var.
//
// Ejemplo: para mapstructure:"MONGO_URI", Viper buscará la clave
// MONGO_URI en el YAML o en la variable de entorno.
type Config struct {
	MongoURI     string   `mapstructure:"MONGO_URI"`
	KafkaBrokers []string `mapstructure:"KAFKA_BROKERS"`
	HttpPort     int      `mapstructure:"HTTP_PORT"`
	OtlpEndpoint string   `mapstructure:"OTLP_ENDPOINT"`
	MetricsAddr  string   `mapstructure:"METRICS_ADDR"`
}

func LoadConfig() (*Config, error) {
	viper.AddConfigPath("config/")
	viper.SetConfigName("config") // sin extensión
	viper.SetConfigType("yaml")
	log := logger.NewLogger()
	viper.AutomaticEnv()

	if err := viper.ReadInConfig(); err != nil {
		log.Error("Error retrieving config file: " + err.Error())
	}

	var cfg Config
	if err := viper.Unmarshal(&cfg); err != nil {
		log.Error("Error parsing config file: " + err.Error())
	}
	return &cfg, nil
}
