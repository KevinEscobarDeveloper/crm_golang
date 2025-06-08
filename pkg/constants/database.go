package constants

import "github.com/spf13/viper"

var (
	DatabaseName   string
	CollectionName string
)

func Init() {
	DatabaseName = viper.GetString("mongo.database")
	CollectionName = viper.GetString("mongo.collection")
}
