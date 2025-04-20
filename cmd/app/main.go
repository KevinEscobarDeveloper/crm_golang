package main

import (
	"go.mongodb.org/mongo-driver/v2/mongo"
	"go.uber.org/fx"
	"mango_crm/pkg/config"
	"mango_crm/pkg/db"
	"mango_crm/pkg/logger"
)

var appModule = fx.Options(
	fx.Provide(logger.NewLogger, config.LoadConfig, db.NewMongoClient),
	fx.Invoke(run),
)

func run(log *logger.Logger, config *config.Config, client *mongo.Client) {
	log.Debug("✅ Logger funcionando desde módulo local.")
	log.Info("Config: " + config.MongoURI)
	log.Info("Config kafka broker: " + config.KafkaBrokers[0])
	log.Info("Orchard application initialized.")
}

func main() {
	app := fx.New(appModule)
	app.Run()
}
