package main

import (
	"context"
	"module/internal/cache"
	"module/internal/handler"
	"module/internal/nats"
	"module/internal/repository/postgres"
	baseService "module/internal/service/base_service"
	"module/server"
	"os"
	"os/signal"
	"syscall"

	_ "github.com/lib/pq"
	"github.com/spf13/viper"

	"github.com/sirupsen/logrus"
)

// @title App API
// @version 1.0
// @description API Server

// @host localhost:8000
// BasePath /
func main() {
	logrus.SetFormatter(new(logrus.JSONFormatter))

	if err := InitConfig(); err != nil {
		logrus.Fatalf("ERROR initializing configs:%s", err.Error())
	}

	db, err := postgres.NewPgSQLConnection(postgres.Config{
		Host:     viper.GetString("db.host"),
		Port:     viper.GetString("db.port"),
		Username: viper.GetString("db.username"),
		Password: os.Getenv("DB_PSWD"),
		DBName:   viper.GetString("db.dbname"),
		SSLMode:  viper.GetString("db.sslmode"),
	})
	if err != nil {
		logrus.Fatalf("ERROR connecting db:%s", err)
	}
	defer func() {
		if err := db.Close(); err != nil {
			logrus.Fatalf("ERROR shutting down db: %s", err.Error())
		}
	}()
	natsConf := nats.NatsConfig{
		ClusterId: viper.GetString("nats.cluster_id"),
		ClientId:  viper.GetString("nats.client_id"),
		Host:      viper.GetString("nats.host"),
		Port:      viper.GetString("nats.port"),
	}
	natsConn, err := nats.NewNatsConnect(natsConf)
	if err != nil {
		logrus.Fatalf("ERROR connecting NATS:%s", err.Error())
	}
	defer func() {
		if err := natsConn.Close(); err != nil {
			logrus.Fatalf("ERROR shutting down NATS:%s", err.Error())
		}
	}()
	repos := postgres.NewPGRepository(db)
	services := baseService.NewBaseService(repos)
	cache := cache.NewCache()
	handler := handler.NewHandler(services, natsConn, cache)
	natsHandler := nats.NewNatsHandler(natsConn, natsConf, services)
	subs, err := natsHandler.InitSubs()
	if err != nil {
		logrus.Fatalf("ERROR subscribing NATS:%s", err.Error())
	}
	defer subs.Unsubscribe()

	srv := new(server.Server)
	go func() {
		if err := srv.Run(viper.GetString("server.port"), handler.InitRoutes()); err != nil {
			logrus.Fatalf("ERROR running server:%s", err.Error())
		}
	}()
	defer func() {
		if err := srv.Shutdown(context.Background()); err != nil {
			logrus.Fatalf("ERROR shutting down server: %s", err.Error())
		}
	}()

	logrus.Print("App started")
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGTERM, syscall.SIGINT)
	<-quit

	logrus.Print("App Shutting Down")
}

func InitConfig() error {
	viper.AddConfigPath("configs")
	viper.SetConfigName(os.Getenv("CONFIG_FILE"))
	return viper.ReadInConfig()
}
