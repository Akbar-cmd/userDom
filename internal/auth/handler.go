package auth

import (
	"github.com/gin-gonic/gin"
	"net/http"

	"userDomain/internal/config"
	"userDomain/pkg/jwt"
	"userDomain/pkg/req"
)

type AuthHandlerDeps struct {
	*config.Config
	*AuthService
}

type AuthHandler struct {
	*config.Config
	*AuthService
}

func NewAuthHandler(router *gin.Engine, deps AuthHandlerDeps) {
	handler := &AuthHandler{
		Config:      deps.Config,
		AuthService: deps.AuthService,
	}

	authGroup := router.Group("/auth")
	{
		authGroup.POST("/login", handler.Login)
		authGroup.POST("/register", handler.Register)
	}
}

func (handler *AuthHandler) Login(c *gin.Context) {

	var body LoginRequest

	err := c.ShouldBindJSON(&body)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Неправильное тело запрсоа"})
		return
	}
	email, err := handler.AuthService.Login(
		body.Email,
		body.Password,
	)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{
			"error": err.Error(),
		})
		return
	}

	token, err := jwt.NewJWT(handler.Config.Token).Create(jwt.JWTData{
		Email: email,
	})
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
	}

	err = req.IsValid(&body)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Введите корректные данные"})
		return
	}

	data := LoginResponse{
		Token: token,
	}
	c.JSON(http.StatusOK, data)
}

func (handler *AuthHandler) Register(c *gin.Context) {
	var body RegisterRequest

	err := c.ShouldBindJSON(&body)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Неправильное тело запроса"})
		return
	}

	err = req.IsValid(&body)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Введите корректные данные"})
		return
	}

	email, err := handler.AuthService.Register(
		body.Email,
		body.Password,
		body.Name,
		body.Surname,
	)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{
			"error": err.Error(),
		})
		return
	}

	token, err := jwt.NewJWT(handler.Config.Token).Create(jwt.JWTData{
		Email: email,
	})
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
	}

	data := RegisterResponse{
		Token: token,
	}

	c.JSON(201, data)
}
