package app

import (
	"context"
	"database/sql"
	"fmt"
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/Yavuzlar/CodinLab/internal/config"
	"github.com/Yavuzlar/CodinLab/internal/http"
	"github.com/Yavuzlar/CodinLab/internal/http/middlewares"
	"github.com/Yavuzlar/CodinLab/internal/http/response"
	"github.com/Yavuzlar/CodinLab/internal/http/server"
	"github.com/Yavuzlar/CodinLab/internal/repositories"
	"github.com/Yavuzlar/CodinLab/internal/services"
	dbadapters "github.com/Yavuzlar/CodinLab/pkg/db_adapters"
	hasher_service "github.com/Yavuzlar/CodinLab/pkg/hasher"
	"github.com/Yavuzlar/CodinLab/pkg/validator_service"
	"github.com/google/uuid"
	"github.com/pressly/goose/v3"
)

func Run(cfg *config.Config) {
	//postgreClient
	conn, err := dbadapters.NewSqlite("./data.db") // sqlite3 için database oluşturma
	if err != nil {
		panic(err) // hata olursa programı durdur
	}
	// database migrate
	err = databaseMigrate(cfg.Application.MigrationsPath, conn.DB)
	if err != nil {
		panic(err)
	}
	// repository initialize
	userRepo := repositories.NewUserRepository(conn)
	logRepo := repositories.NewLogRepository(conn)

	// utilities initialize
	validator := validator_service.NewValidatorService()

	// service initialize

	allService := services.CreateNewServices(userRepo, logRepo, validator)

	//---------------first run -----------------
	hashedPass, err := hasher_service.HashPassword(cfg.Application.Managment.ManagmentPassword)
	if err != nil {
		panic(err)
	}
	firstRun(conn.DB, cfg.Application.Managment.ManagmentUsername, hashedPass)
	//--------------------------------------------

	//handler initialize
	handlers := http.NewHandler(allService)

	//server initialize
	fiberSrv := server.NewServer(cfg, response.ResponseHandler)

	//captcha store initialize

	go func() {
		err := fiberSrv.Run(handlers.Init(cfg.Application.DevMode, middlewares.InitMiddlewares(cfg)...))
		if err != nil {
			log.Fatalf("Error while running fiber server: %v", err.Error())
		}
	}()
	c := make(chan os.Signal, 1)                    // Create channel to signify a signal being sent
	signal.Notify(c, os.Interrupt, syscall.SIGTERM) // When an interrupt or termination signal is sent, notify the channel
	<-c                                             // This blocks the main thread until an interrupt is received
	fmt.Println("Gracefully shutting down...")      // Daha iyi yapılabilir !!
	_ = fiberSrv.Shutdown(context.Background())
	if err := conn.Close(); err != nil {
		fmt.Println("Error while closing database connection: ", err.Error())
	}
	fmt.Println("Fiber was successful shutdown.")
}

func databaseMigrate(migrationPath string, db *sql.DB) error {
	err := goose.SetDialect("sqlite3")
	if err != nil {
		return err
	}
	if err := goose.Up(db, migrationPath); err != nil {
		return err
	}
	return nil
}

func firstRun(conn *sql.DB, username, pass string) {
	tx, err := conn.Begin()
	if err != nil {
		panic(err)
	}
	varmi := new(int8)
	err = tx.QueryRow("SELECT COUNT(*) FROM t_users WHERE role = 'admin'").Scan(varmi)
	if err != nil && err != sql.ErrNoRows {
		panic(err)
	}
	if *varmi == 0 {
		uuid := uuid.New()
		stm, err := tx.Prepare("INSERT INTO t_users (id, username, name, surname, password, role) VALUES (?, ?, ?, ?, ?, ?)")
		if err != nil {
			panic(err)
		}
		res, err := stm.Exec(uuid.String(), username, "admin", "admin", pass, "admin")
		if err != nil {
			panic(err)
		}
		_, err = res.LastInsertId()
		if err != nil {
			panic(err)
		}
		err = tx.Commit()
		if err != nil {
			panic(err)
		}
	}
}
