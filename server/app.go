package server

import (
	"chatapp/auth"
	"chatapp/chat"
	"chatapp/rooms"
	//authdelivery "chatapp/auth/delivery"
	authUsecase "chatapp/auth/usecase"
	authrepos "chatapp/auth/repository"
	chatdelivery "chatapp/chat/delivery"
	chatUsecase "chatapp/chat/usecase"
	chatrepos "chatapp/chat/repository"
	roomsdelivery "chatapp/rooms/delivery"
	roomsUsecase "chatapp/rooms/usecase"
	roomsrepos "chatapp/rooms/repository"

	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"

	"github.com/gin-gonic/gin"
	_ "github.com/jackc/pgx"
	"github.com/jmoiron/sqlx"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type App struct{
	server *http.Server
	authUC auth.UseCase
	chatUC chat.UseCase
	roomsUC rooms.UseCase
}

func NewApp() *App{
	postgresDB := initPostgreDB()
	authRepos := authrepos.NewUserRepository(postgresDB)

	mongoDB := initMongoDB()
	chatRepos := chatrepos.NewChatRepository(mongoDB)

	roomsRepos := roomsrepos.NewRoomRepository(mongoDB, "collection-name")

	return &App{
		authUC: authUsecase.NewAuthUseCase(authRepos, "xcdPO78_$hq", []byte("xpasretvbn"), 10000),
		chatUC: chatUsecase.NewChatUseCase(chatRepos),
		roomsUC: roomsUsecase.NewRoomsUseCase(roomsRepos),
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
	
	db, err := sqlx.Open("pgx", ReadPostgresConfigs())
	if err != nil {
		log.Fatal("init postgres: ", err)
	}

	return db
}

func (a *App) Run() error{
	router := gin.Default()

	router.StaticFS("/static/", http.Dir("./client/templates/chat/static/"))

	api := router.Group("/api" /*authMiddleware.Handle*/ )

	chatHandler := chatdelivery.NewHandler(a.chatUC)
	roomsHandler := roomsdelivery.NewHandler(a.roomsUC)

	chats := api.Group("my-chats")
	chats.GET("/", roomsHandler.GetAllRoomsList)
	chats.GET("/:chat_id/info", roomsHandler.GetRoom)
	chats.GET("/:chat_id", chatHandler.WSEndpoint)
	chats.GET("/:chat_id/participants", /*GetParticipantsFunc*/)
	chats.PUT("/:chat_id/participants/add", /*AddParticipantsFunc*/)


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