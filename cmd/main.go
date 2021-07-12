package main

import (
	"context"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
	"github.com/p12s/2gis-catalog-api"
	"github.com/p12s/2gis-catalog-api/pkg/handler"
	"github.com/p12s/2gis-catalog-api/pkg/repository"
	"github.com/p12s/2gis-catalog-api/pkg/service"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	"os"
	"os/signal"
	"syscall"
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

	db, err := repository.NewPostgresDB(repository.Config{
		Driver:   viper.GetString("db.driver"),
		Host:     viper.GetString("db.host"),
		Port:     viper.GetString("db.port"),
		Username: viper.GetString("db.username"),
		Password: os.Getenv("DB_PASSWORD"),
		DBName:   viper.GetString("db.dbname"),
		SSLMode:  viper.GetString("db.sslmode"),
	})
	if err != nil {
		logrus.Fatalf("👺👺👺 failed to initialize db: %s\n", err.Error())
	}

	// инит клин (репо-сервис-хендлер)
	repos := repository.NewRepository(db)
	services := service.NewService(repos)
	handlers := handler.NewHandler(services)

	// ран сервер
	srv := new(common.Server)
	go func() {
		if err := srv.Run(viper.GetString("port"), handlers.InitRoutes()); err != nil {
			logrus.Fatalf("error while running http server: %s\n", err.Error())
		}
	}()
	logrus.Print("😀😀😀 app started 😀😀😀")

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGTERM, syscall.SIGINT)
	<-quit

	logrus.Print("🧟🧟🧟 TodoApp Shutting Down 🧟🧟🧟")

	if err := srv.Shutdown(context.Background()); err != nil {
		logrus.Errorf("error occured on server shutting down: %s", err.Error())
	}

	if err := db.Close(); err != nil {
		logrus.Errorf("error occured on db connection close: %s", err.Error())
	}
}

func initConfig() error {
	viper.AddConfigPath("configs")
	viper.SetConfigName("config")
	return viper.ReadInConfig()
}
