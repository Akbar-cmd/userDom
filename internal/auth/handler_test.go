package auth_test

import (
	"bytes"
	"encoding/json"
	"github.com/DATA-DOG/go-sqlmock"
	"github.com/gin-gonic/gin"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"net/http"
	"net/http/httptest"
	"testing"

	"userDomain/internal/auth"
	"userDomain/internal/config"
	"userDomain/internal/user"
	db "userDomain/pkg/storage/postgres"
)

func bootstrap() (*auth.AuthHandler, sqlmock.Sqlmock, error) { // создаем моковую бд и подключаем
	database, mock, err := sqlmock.New()
	if err != nil {
		return nil, nil, err
	}
	gormDb, err := gorm.Open(postgres.New(postgres.Config{
		Conn: database,
	}))
	if err != nil {
		return nil, nil, err
	}

	// передаем мок бд в репу
	userRepo := user.NewUserRepository(&db.Db{
		gormDb,
	})

	// Реализуем хэндлер
	handler := auth.AuthHandler{
		Config:      &config.Config{Token: "token"},
		AuthService: auth.NewAuthService(userRepo),
	}
	return &handler, mock, nil
}

func TestLoginHandlerSuccess(t *testing.T) {
	// подготавливаем запрос
	handler, mock, err := bootstrap()
	// добавляем новые колонки
	rows := sqlmock.NewRows([]string{"email", "password"}).
		// добавляем значения созданных колонок
		AddRow("a2@a.ru", "$2a$10$qux9qZwaqoGQ3a00AOyLfeqLitkLJBIYJC/kOwAj.Rbxg6bWitl7K")
	// Если в нашей БД кто-то будет делать select, то вернет колонки
	mock.ExpectQuery("SELECT").WillReturnRows(rows)

	// мокаем Insert
	// начинаем транзакцию
	mock.ExpectBegin()
	// выполняем

	// коммитим
	mock.ExpectCommit()
	if err != nil {
		t.Fatal(err)
		return
	}

	data, _ := json.Marshal(&auth.LoginRequest{
		Email:    "a2@a.ru",
		Password: "1",
	})
	reader := bytes.NewReader(data)

	// то, что будет записывать ответ
	w := httptest.NewRecorder()
	// то, что будет передавать
	req := httptest.NewRequest(http.MethodPost, "/auth/login", reader)

	// Создаем gin
	c, _ := gin.CreateTestContext(w)
	c.Request = req

	// передаем в Login
	handler.Login(c)

	if w.Code != http.StatusOK {
		t.Errorf("got %d, expected %d", w.Code, 201)
	}
}

func TestRegisterHandlerSuccess(t *testing.T) {
	// подготавливаем запрос
	handler, mock, err := bootstrap()
	// добавляем новые колонки
	rows := sqlmock.NewRows([]string{"email", "password", "name", "surname"})
	// Если в нашей БД кто-то будет делать select, то вернет колонки
	mock.ExpectQuery("SELECT").WillReturnRows(rows)

	// мокаем Insert
	// начинаем транзакцию
	mock.ExpectBegin()
	// выполняем Возвращаем 1 как новый созданный id
	mock.ExpectQuery("INSERT").WillReturnRows(sqlmock.NewRows([]string{"id"}).AddRow(1))
	// коммитим
	mock.ExpectCommit()
	if err != nil {
		t.Fatal(err)
		return
	}

	data, _ := json.Marshal(&auth.RegisterRequest{
		Email:    "a2@a.ru",
		Password: "1",
		Name:     "Вася",
		Surname:  "Пупкин",
	})
	reader := bytes.NewReader(data)

	// то, что будет записывать ответ
	w := httptest.NewRecorder()
	// то, что будет передавать
	req := httptest.NewRequest(http.MethodPost, "/auth/register", reader)

	// Создаем gin
	c, _ := gin.CreateTestContext(w)
	c.Request = req

	// передаем в register
	handler.Register(c)

	if w.Code != http.StatusCreated {
		t.Errorf("got %d, expected %d", w.Code, 201)
	}
}
