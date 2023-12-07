package main

import (
	"context"
	"github.com/redis/go-redis/v9"
	"github.com/spf13/viper"
	"log"
	"main_terminal"
	"main_terminal/internal/handler"
	"main_terminal/internal/infrastructure"
	"main_terminal/internal/service"
	"os"
	"os/signal"
	"syscall"
)

func main() {

	if err := InitConfig(); err != nil {
		log.Fatalf("Error while initializating configs: %s", err.Error())
	}

	CConfig := &redis.Options{
		Addr:     viper.GetString("cache.redis.Addr"),
		Password: viper.GetString("cache.redis.Password"),
		DB:       viper.GetInt("cache.redis.DB"),
	}
	StocksAPI := viper.GetString("exAPI.stocksAPI")
	CurrencyAPI := viper.GetString("exAPI.currencyAPI")

	infrastructures := infrastructure.NewInfrastructure(
		nil,
		CConfig,
		StocksAPI,
		CurrencyAPI,
	)

	services := service.NewService(infrastructures)
	handlers := handler.NewHandler(services)

	srv := new(main_terminal.Server)
	go func() {
		if err := srv.Run(viper.GetString("server.port"), handlers.GinHttp.InitRoutes()); err != nil {
			log.Fatalf("error occured while initializing http server: %s", err.Error())
		}
	}()

	log.Println("Main terminal was started successfully")

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGTERM, syscall.SIGINT)
	<-quit

	log.Println("Main terminal shutting down")

	if err := srv.Shutdown(context.Background()); err != nil {
		log.Fatalf("Error while shutting down server: %s", err.Error())
	}
}

func InitConfig() error {
	viper.AddConfigPath("configs")
	viper.SetConfigName("config")
	return viper.ReadInConfig()
}
