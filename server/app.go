package server

import (
	"chatapp/accounts"
	"chatapp/auth"
	"chatapp/chat"
	"chatapp/rooms"

	"fmt"

	authdelivery "chatapp/auth/delivery"
	authrepos "chatapp/auth/repository"
	authUsecase "chatapp/auth/usecase"

	chatdelivery "chatapp/chat/delivery"
	chatrepos "chatapp/chat/repository"
	chatUsecase "chatapp/chat/usecase"

	roomsdelivery "chatapp/rooms/delivery"
	roomsrepos "chatapp/rooms/repository"
	roomsUsecase "chatapp/rooms/usecase"

	accountsdelivery "chatapp/accounts/delivery"
	accountsrepos "chatapp/accounts/repository"
	accountsUsecase "chatapp/accounts/usecase"

	"chatapp/app/clienthandler"

	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"

	"github.com/gin-gonic/gin"
	_ "github.com/jackc/pgx"
	"github.com/jmoiron/sqlx"
	"github.com/joho/godotenv"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

type App struct {
	server     *http.Server
	authUC     auth.UseCase
	chatUC     chat.UseCase
	roomsUC    rooms.UseCase
	accountsUC accounts.UseCase
}

func NewApp() *App {
	postgresDB := initPostgreDB()
	mongoDB := initMongoDB()

	chatRepos := chatrepos.NewChatRepository(mongoDB)
	authRepos := authrepos.NewUserRepository(postgresDB)
	roomsRepos := roomsrepos.NewRoomRepository(mongoDB, postgresDB)
	accountsRepos := accountsrepos.NewAccountsRepository(postgresDB)

	godotenv.Load("postgres.env")
	key := os.Getenv("SIGNING_KEY")
	tokenTLS := 24 * 60 * 60 * time.Second

	fmt.Println("KEY : ", key)

	return &App{
		authUC:     authUsecase.NewAuthUseCase(authRepos, os.Getenv("HASH_SALT"), key, tokenTLS),
		chatUC:     chatUsecase.NewChatUseCase(chatRepos),
		roomsUC:    roomsUsecase.NewRoomsUseCase(roomsRepos),
		accountsUC: accountsUsecase.NewAccountsUseCase(accountsRepos),
	}
}

func initMongoDB() *mongo.Database {

	godotenv.Load("mongo.env")

	mongoClient, err := mongo.NewClient(options.Client().ApplyURI(os.Getenv("MONGO_URI")))
	if err != nil {
		log.Fatal(err)
	}
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	err = mongoClient.Connect(ctx)
	if err != nil {
		log.Fatal(err)
	}

	if err := mongoClient.Ping(context.TODO(), readpref.Primary()); err != nil {
		log.Fatal(err)
	}

	db := mongoClient.Database(os.Getenv("MONGO_DB"))
	return db
}

func initPostgreDB() *sqlx.DB {

	db, err := sqlx.Open("pgx", ReadPostgresConfigs())
	if err != nil {
		log.Fatal("init postgres: ", err)
	}

	return db
}

func (a *App) Run() error {
	router := gin.Default()

	//auth-api endpoints
	authHandler := authdelivery.NewHadler(a.authUC)

	router.POST("/sign-up", authHandler.SignUp)
	router.POST("/sign-in", authHandler.SignIn)

	authMiddleware := authdelivery.NewAuthMiddleware(a.authUC)
	api := router.Group("/api", authMiddleware.Handle)

	chatHandler := chatdelivery.NewHandler(a.chatUC, a.authUC)
	roomsHandler := roomsdelivery.NewHandler(a.roomsUC, a.authUC)
	accountsHandler := accountsdelivery.NewHandler(a.accountsUC)

	api.POST("/create-chat/common", roomsHandler.CreateCommonRoom)
	api.POST("/create-chat/group", roomsHandler.CreateGroupRoom)

	chats := api.Group("my-chats")
	chats.GET("/", roomsHandler.GetAllRoomsList)
	chats.GET("/:chat_id/info", roomsHandler.GetRoom)
	chats.GET("/:chat_id", chatHandler.HandleConnections)
	go chatHandler.HandleMessages()
	chats.GET("/:chat_id/participants" /*GetParticipantsFunc*/)
	chats.POST("/:chat_id/participants/add", roomsHandler.AddParticipants)

	accounts := api.Group("accounts")
	accounts.GET("/find/:user", accountsHandler.FindUser)

	//client browser pages

	//serving static files
	router.StaticFS("/static/", http.Dir("./app/templates/"))

	router.GET("sign-up", clienthandler.SignUpPage)
	router.GET("sign-in", clienthandler.SignInPage)

	app := router.Group("app", authMiddleware.Handle, clienthandler.CacheContolMiddleware)
	app.GET("/chats", clienthandler.ChatListsPage)
	app.GET("/chats/:chat_id", clienthandler.ChatPage)

	a.server = &http.Server{
		Addr:           ":8090",
		MaxHeaderBytes: 1 << 20,
		Handler:        router,
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
