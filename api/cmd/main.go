package main

import (
	"github.com/ivanovamir/grpc-besm/api/internal/handler"
	"github.com/ivanovamir/grpc-besm/api/internal/repository"
	"github.com/ivanovamir/grpc-besm/api/internal/service"
	"github.com/ivanovamir/grpc-besm/api/server"
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/joho/godotenv"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

func main() {
	logrus.SetFormatter(&logrus.JSONFormatter{})

	if err := godotenv.Load(".env"); err != nil {
		logrus.Fatalf("error loading env variables: %s", err.Error())
	}

	if err := initConfig(); err != nil {
		log.Fatalf("error initializing configs: %s", err.Error())
	}
	db, err := repository.NewPostgresDB(repository.Config{
		Host:     viper.GetString("db.host"),
		Port:     viper.GetString("db.port"),
		Username: viper.GetString("db.username"),
		Password: os.Getenv("DB_PASSWORD"),
		DBName:   viper.GetString("db.dbname"),
	})

	NewRepos := repository.NewRepository(db)
	NewService := service.NewService(NewRepos)
	NewHandler := handler.NewHandler(NewService)

	if err != nil {
		logrus.Fatalf("error initializing db: %s", err.Error())
	}

	grpcSrv := new(server.GrpcServer).NewServer(NewHandler)

	go func() {
		logrus.Print("gRPC server started")
		if err := grpcSrv.ListenAndServe(viper.GetString("srv.conn"), viper.GetString("srv.port")); err != nil {
			logrus.Print("gRPC server error: %s", err.Error())
		}
	}()

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGTERM, syscall.SIGINT)

	<-quit

	logrus.Print("gRPC app shutting down")

	grpcSrv.Stop()

	sqlDB, err := db.DB()
	if err := sqlDB.Close(); err != nil {
		logrus.Errorf("error occured on db connection close: %s", err.Error())
	}
}

func initConfig() error {
	viper.AddConfigPath("api/configs")
	viper.SetConfigName("config")
	return viper.ReadInConfig()
}
