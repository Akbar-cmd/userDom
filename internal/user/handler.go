package user

import (
	"context"
	"fmt"
	"github.com/gin-gonic/gin"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"log"
	"net/http"
	"strconv"
	"time"

	"userDomain/internal/config"
	desc "userDomain/pkg/gRPC"
	"userDomain/pkg/middleware"
)

type UserHandlerDeps struct {
	*config.Config
	BikeClient desc.BikeClient
}

type UserHandler struct {
	*config.Config
	BikeClient desc.BikeClient
}

var body struct {
	Model  string `json:"model" binding:"required"`
	Color  string `json:"color" binding:"required"`
	IsWork bool   `json:"is_work"`
}

func NewUserHandler(router *gin.Engine, deps UserHandlerDeps) {
	handler := &UserHandler{
		Config:     deps.Config,
		BikeClient: deps.BikeClient,
	}

	userGroup := router.Group("/user")
	{
		userGroup.GET("/:id", middleware.IsAuthed(deps.Config), handler.GetAllBikes)
		userGroup.POST("/:id", middleware.IsAuthed(deps.Config), handler.CreateBike)
		userGroup.PATCH("/:id/:bike_id", middleware.IsAuthed(deps.Config), handler.UpdateBike)
		userGroup.DELETE("/:id/:bike_id", middleware.IsAuthed(deps.Config), handler.DeleteBike)
	}
}

// Handler для получения велосипеда через gRPC
func (handler *UserHandler) GetAllBikes(c *gin.Context) {
	// Вытаскиваем данные из контекста
	email, ok := c.Request.Context().Value(middleware.ContextEmailKey).(string)
	if ok {
		fmt.Println(email)
	}

	userID, err := getUserID(c)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
		return
	}

	// Создаем контекст с таймаутом для gRPC запроса
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	// Делаем gRPC запрос
	resp, err := handler.BikeClient.GetBikeByUser(ctx, &desc.GetBikeByUserRequest{
		UserId: userID,
	})
	if err != nil {
		log.Printf("Failed to get bike by user ID: %v", err)
		c.JSON(500, gin.H{"error": "Не удалось получить информацию о велосипеде"})
		return
	}

	// Преобразование в удобный для клиента формат
	bikes := make([]map[string]interface{}, 0, len(resp.GetBikes()))
	for _, bike := range resp.GetBikes() {
		bikes = append(bikes, map[string]interface{}{
			"bike_id": bike.GetBikeId(),
			"model":   bike.GetModel(),
			"color":   bike.GetColor(),
			"is_work": bike.GetIsWork(),
		})
	}

	// Возвращаем результат в JSON формате
	c.JSON(200, gin.H{
		"user_id": userID,
		"bike":    bikes,
	})

}

func (handler *UserHandler) CreateBike(c *gin.Context) {

	email, ok := c.Request.Context().Value(middleware.ContextEmailKey).(string)
	if ok {
		fmt.Println(email)
	}

	userID, err := getUserID(c)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
		return
	}

	err = c.ShouldBindJSON(&body)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Неправильный формат данных"})
		return
	}

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	resp, err := handler.BikeClient.CreateBikeByUser(ctx, &desc.CreateBikeByUserRequest{
		UserId: userID,
		Model:  body.Model,
		Color:  body.Color,
		IsWork: body.IsWork,
	})

	if err != nil {
		handleGRPCError(c, err)
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"bike":    resp.GetBikeId(),
		"message": "Велосипед успешно создан",
	})
}

func (handler *UserHandler) UpdateBike(c *gin.Context) {

	userID, err := getUserID(c)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Введите ID"})
		return
	}

	bikeIDStr := c.Param("bike_id")
	bikeID, err := strconv.ParseUint(bikeIDStr, 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid bike ID"})
		return
	}

	err = c.ShouldBindJSON(&body)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Неправильный формат данных"})
		return
	}

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	_, err = handler.BikeClient.UpdateBikeByUser(ctx, &desc.UpdateBikeByUserRequest{
		UserId: userID,
		BikeId: bikeID,
		Model:  body.Model,
		Color:  body.Color,
		IsWork: body.IsWork,
	})

	if err != nil {
		handleGRPCError(c, err)
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"message": "Данные о велосипеде успешно обновлены",
	})

}

func (handler *UserHandler) DeleteBike(c *gin.Context) {

	userID, err := getUserID(c)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Введите ID"})
		return
	}

	bikeIDStr := c.Param("bike_id")
	bikeID, err := strconv.ParseUint(bikeIDStr, 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid bike ID"})
		return
	}

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	_, err = handler.BikeClient.DeleteBikeByUser(ctx, &desc.DeleteBikeByUserRequest{
		UserId: userID,
		BikeId: bikeID,
	})

	if err != nil {
		handleGRPCError(c, err)
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"message": "Данные о велосипеде успешно удалены",
	})

}

func getUserID(c *gin.Context) (uint64, error) {
	userIDStr := c.Param("id")
	if userIDStr == "" {
		return 0, status.Error(codes.Unauthenticated, "User не авторизован")

	}
	userID, err := strconv.ParseUint(userIDStr, 10, 32)
	if err != nil {
		return 0, status.Error(codes.Unauthenticated, "User не авторизован")
	}
	return userID, nil

}

func handleGRPCError(c *gin.Context, err error) {
	st, ok := status.FromError(err)
	if !ok {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Unknown error"})
		return
	}

	switch st.Code() {
	case codes.InvalidArgument:
		c.JSON(http.StatusBadRequest, gin.H{"error": st.Message()})
	case codes.NotFound:
		c.JSON(http.StatusNotFound, gin.H{"error": st.Message()})
	case codes.Unauthenticated:
		c.JSON(http.StatusUnauthorized, gin.H{"error": st.Message()})
	default:
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Internal server error"})
	}
}
