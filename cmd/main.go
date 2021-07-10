package main

import (
	"github.com/joho/godotenv"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

/*
company := &common.Company{
		Id: 1,
		Name: "Yandex",
		Phones: []string{"22-22-22", "33-33-33"},
		Building: common.Building{
			Address: common.Address{
				City: "Moscow",
				Street: "Lenina",
				House: 1,
			},
			Coordinates: common.Coordinates{
				Longitude: 38.8951,
				Latitude: -77.0364,
			},
		},
		Rubric: []common.Rubric{
			{
				Name: "Интернет-услуги",
			},
			{
				Name: "Услуги доставки",
			},
			{
				Name: "Облачные услуги",
			},
			{
				Name: "Транспортные услуги",
			},
			{
				Name: "Финансовые услуги",
			},
		},
	}
 */

// @title Catalog App API
// @version 0.0.1
// @description API Server for company catalog
// @host localhost:80
// @BasePath /
func main() {
	logrus.SetFormatter(new(logrus.JSONFormatter))
	if err := initConfig(); err != nil {
		logrus.Fatalf("👺👺👺 error init configs/config.yml file: %s", err.Error())
	}
	if err := godotenv.Load(); err != nil {
		logrus.Fatalf("👺👺👺 error load .env file variables: %s", err.Error())
	}

}

func initConfig() error {
	viper.AddConfigPath("configs")
	viper.SetConfigName("config")
	return viper.ReadInConfig()
}