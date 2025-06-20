package middleware

import (
	"context"
	"github.com/gin-gonic/gin"
	"net/http"
	"strings"

	"userDomain/internal/config"
	"userDomain/pkg/jwt"
)

type key string

const (
	// заготовили ключ по которому добавить в ctx
	ContextEmailKey key = "ContextEmailKey"
)

func writeUnauthed(c *gin.Context) {
	c.AbortWithStatus(http.StatusUnauthorized)
	c.String(http.StatusUnauthorized, http.StatusText(http.StatusUnauthorized))
}

func IsAuthed(config *config.Config) gin.HandlerFunc {
	return func(c *gin.Context) {
		// Получаем токен из заголовка
		authedHeader := c.GetHeader("Authorization")
		// Проверяем начинается ли с "Bearer "
		if !strings.HasPrefix(authedHeader, "Bearer ") {
			writeUnauthed(c)
			return
		}

		// Извлекаем токен
		token := strings.TrimPrefix(authedHeader, "Bearer ")

		// Парсим
		isValid, data := jwt.NewJWT(config.Token).Parse(token)
		if !isValid {
			writeUnauthed(c)
			return
		}

		// Передаем данные в ctx
		// создали новый context на базе существующего добавив ключ и значение
		ctx := context.WithValue(c.Request.Context(), ContextEmailKey, data.Email)
		// обновляем информацию
		c.Request = c.Request.WithContext(ctx)
		c.Next()
	}
}
