package main

import (
	"log"

	"todo/pkg/handler"
	"todo/pkg/repository"
	"todo/pkg/service"

	"todo"

	"github.com/spf13/viper"
)

func main() {
	if err := initConfig(); err != nil {
		log.Fatalf("error initializing config: %s", err.Error())
	}

	db, err := repository.NewSqliteDB(&repository.Config{
		Driver: viper.GetString("db.driver"),
		Dsn:    viper.GetString("db.dsn"),
	})
	if err != nil {
		log.Fatalf("failed to initialize db: %s", err.Error())
	}

	repos := repository.NewRepository(db)
	services := service.NewService(repos)
	handlers := handler.NewHandler(services)
	svr := new(todo.Server)
	if err := svr.Run(viper.GetString("port"), handlers.InitRoutes()); err != nil {
		log.Fatalf("error occured while running http server: %s", err.Error())
	}
}

func initConfig() error {
	viper.AddConfigPath("configs")
	viper.SetConfigName("config")
	return viper.ReadInConfig()
}
