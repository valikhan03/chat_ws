package server

import (
	"chatapp/auth"
	"chatapp/auth/delivery/http"
	"chatapp/auth/repository/authdatabase"
	authUsecase "chatapp/auth/usecase"
	"chatapp/chat"
	"chatapp/chat/delivery"
	"chatapp/chat/repository"
	chatUsecase "chatapp/chat/usecase"
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/jmoiron/sqlx"
	_ "github.com/jackc/pgx/stdlib"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	
)

type App struct{
	server *http.Server
	authUC auth.UseCase
	chatUC chat.UseCase
}

func NewApp() *App{
	postgresDB := initPostgreDB()
	authRepos := authdatabase.NewUserRepository(postgresDB)

	mongoDB := initMongoDB()
	chatRepos := chatdatabase.NewChatRepository(mongoDB)

	return &App{
		authUC: authUsecase.NewAuthUseCase(authRepos, "xcdPO78_$hq", []byte("xpasretvbn"), 10000),
		chatUC: chatUsecase.NewChatUseCase(chatRepos),
	}
}

func initMongoDB() *mongo.Database{
	configs := ReadMongoConfigs()
	mongo, err := mongo.NewClient(options.Client().ApplyURI(configs))
	if err != nil {
		log.Fatal(err)
	}
	db := mongo.Database("")

	return db
}

func initPostgreDB() *sqlx.DB {
	db, err := sqlx.Connect("postgres", ReadPostgresConfigs())
	if err != nil {
		log.Fatal("init postgres: ", err)
	}

	return db
}

func (a *App) Run() error{
	router := gin.Default()

	router.StaticFS("/static/", http.Dir("./client/templates/chat/static/"))

	
	wsdelivery.RegisterChatHTTPWSEndpoints(router, a.chatUC)
	authhttp.RegisterAuthHTTPEndpoints(router, a.authUC)

	a.server = &http.Server{
		Addr: ":8090",
		MaxHeaderBytes: 1 << 20,
		Handler: router,
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