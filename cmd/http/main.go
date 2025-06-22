package main

import (
	"context"
	"fmt"
	"github.com/gin-gonic/gin"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/grpc/status"
	"log"
	"os"
	"time"

	"userDomain/internal/auth"
	"userDomain/internal/config"
	"userDomain/internal/metric"
	"userDomain/internal/user"
	desc "userDomain/pkg/gRPC"
	"userDomain/pkg/storage/postgres"
	"github.com/zsais/go-gin-prometheus"
)

const (
	envLocal = "local"
	envDev   = "dev"
	envProd  = "prod"
)

const (
	address = "grpc-server:50051"
)

// Глобальный gRPC клиент
var grpcClient desc.BikeClient
var grpcConn *grpc.ClientConn

// Функция для инициализации gRPC соединения
func initGRPCClient() error {

	var conn *grpc.ClientConn
	var err error

	for i := 0; i < 5; i++ {
		conn, err = grpc.NewClient(
			address,
			grpc.WithTransportCredentials(insecure.NewCredentials()),
			grpc.WithBlock(), // Ждем установления соединения
			grpc.WithUnaryInterceptor(metric.ClientMetricsInterceptor()),
		)
		if err == nil {
			break
		}
		log.Printf("Attempt %d: Failed to connect to gRPC server: %v", i+1, err)
		time.Sleep(2 * time.Second)
	}

	if err != nil {
		return fmt.Errorf("failed to connect to gRPC server after 5 attempts: %v", err)
	}

	grpcClient = desc.NewBikeClient(conn)
	grpcConn = conn
	log.Printf("Successfully connected to gRPC server at %s", address)
	return nil
}

// Проверка состояния gRPC соединения
func checkGRPCConnection() error {
	if grpcConn == nil {
		return fmt.Errorf("gRPC connection is nil")
	}

	// Простой health check через context с таймаутом
	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()

	// Пробуем сделать простой запрос
	_, err := grpcClient.GetBikeByUser(ctx, &desc.GetBikeByUserRequest{UserId: 0})
	if err != nil {
		st, ok := status.FromError(err)
		if ok {
			log.Printf("gRPC health check failed: %s", st.Message())
		}
		return err
	}
	return nil
}

func App() (*gin.Engine, error) {

	// TODO : init client metric
	if err := metric.InitClientMetrics(); err != nil {
		log.Fatal(err)
	}

	// TODO : init gRPC  client
	if err := initGRPCClient(); err != nil {
		log.Fatalf("Ошибка инициализации gRPC клиента: %v", err)
	}

	// Проверяем соединение
	if err := checkGRPCConnection(); err != nil {
		log.Printf("Warning: gRPC connection check failed: %v", err)
	}

	// TODO : init config: cleanenv
	cfg := config.MustLoad()
	fmt.Println(cfg)

	// TODO : init storage: psql
	db, err := postgres.NewDB(cfg.DSN)
	if err != nil {
		log.Fatal("Ошибка подключения к БД: ", err)
		os.Exit(1)
	}

	// TODO : Repo
	userRepository := user.NewUserRepository(db)

	// TODO : Services
	authService := auth.NewAuthService(userRepository)

	// TODO : init router

	// create gin router
	r := gin.Default()

	p := ginprometheus.NewPrometheus("gin")
	p.Use(r)

	r.GET("/metrics", gin.WrapH(promhttp.Handler()))

	// Deps
	auth.NewAuthHandler(r, auth.AuthHandlerDeps{
		Config:      cfg,
		AuthService: authService,
	})

	user.NewUserHandler(r, user.UserHandlerDeps{
		Config:     cfg,
		BikeClient: grpcClient,
	})

	if err != nil {
		return nil, err
	}

	return r, nil
}

func main() {

	router, err := App()
	if err != nil {
		log.Fatal(err)
	}

	log.Print("Сервер запущен на порту 8080")
	if err := router.Run(":8080"); err != nil {
		log.Fatalf("Сбой в работе сервера: %v", err)
	}

}
