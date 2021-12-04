package server

import (
	"chatapp/auth"
	"chatapp/auth/repository/authdatabase"
	authUsecase "chatapp/auth/usecase"
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"

	"github.com/jmoiron/sqlx"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type App struct{
	server *http.Server
	authUC auth.UseCase
}

func NewApp() *App{
	postgresDB := initPostgreDB()

	authRepos := authdatabase.NewUserRepository(postgresDB)

	return &App{
		authUC: authUsecase.NewAuthUseCase(authRepos, "xcdPO78_$hq", []byte("xpasretvbn"), 10000),
	}
}

func initMongoDB() *mongo.Database{
	mongo, err := mongo.NewClient(options.Client().ApplyURI("mongodb://127.0.0.1:27017"))
	if err != nil {
		log.Fatal(err)
	}
	db := mongo.Database("")

	return db
}

func initPostgreDB() *sqlx.DB {
	db, err := sqlx.Connect("", "")
	if err != nil {
		log.Fatal(err)
	}

	return db
}

func (a *App) Run() error{
	a.server = &http.Server{
		Addr: "8090",
		MaxHeaderBytes: 1 << 20,
	}

	go func() {
		if err := a.server.ListenAndServe(); err != nil {
			log.Fatalf("Server failed : %+v", err)
		}
	}()

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt)
	<-quit

	ctx, shutdown := context.WithTimeout(context.Background(), 5*time.Second)
	defer shutdown()

	return a.server.Shutdown(ctx)
}