package main

import (
	"database/sql"
	"fmt"
	"os"
	"time"

	"github.com/labstack/echo"
	_ "github.com/lib/pq"
	"github.com/saeedjalalisj/down-monitor/infra/logger"
	_serviceHttpDelivery "github.com/saeedjalalisj/down-monitor/internal/service/delivery/http"
	_serviceRepo "github.com/saeedjalalisj/down-monitor/internal/service/repository/postgres"
	_serviceUsecase "github.com/saeedjalalisj/down-monitor/internal/service/usecase"
	"go.uber.org/zap"
)

func main() {
	log, err := logger.New("DOWN-MONITOR-API")
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	defer log.Sync()

	// Perform the startup and shutdown sequence.
	if err := run(log); err != nil {
		log.Errorw("startup", "ERROR", err)
		log.Sync()
		os.Exit(1)
	}
}

func run(log *zap.SugaredLogger) error {
	log.Infow("Starting App")
	defer log.Infow("shutdown complete")

	connStr := "user=postgres dbname=postgres password=postgres sslmode=disable"
	conn, err := sql.Open("postgres", connStr)
	if err != nil {
		log.Fatal(err)
	}

	err = conn.Ping()
	if err != nil {
		log.Fatal(err)
	}

	defer func() {
		err := conn.Close()
		if err != nil {
			log.Fatal(err)
		}
	}()
	fmt.Println("Logger program")

	e := echo.New()

	timeoutContext := time.Duration(10) * time.Second

	// service
	serviceRepo := _serviceRepo.NewPostgresServiceRepository(conn)
	su := _serviceUsecase.NewServiceUsecase(serviceRepo, timeoutContext)
	_serviceHttpDelivery.NewServiceHandler(e, su)

	fmt.Println(e.Start(":9090"))

	return nil
}
