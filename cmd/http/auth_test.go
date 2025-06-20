package main

import (
	"bytes"
	"encoding/json"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"

	"userDomain/internal/auth"
)

func TestLoginSuccess(t *testing.T) {
	r, err := App()
	if err != nil {
		t.Fatalf("Failed to initialize app: %v", err)
	}

	// Создаем тестовый	сервер
	ts := httptest.NewServer(r)
	defer ts.Close()

	data, _ := json.Marshal(&auth.LoginRequest{
		Email:    "a2@a.ru",
		Password: "1",
	})

	res, err := http.Post(ts.URL+"/auth/login", "application/json", bytes.NewReader(data))
	if err != nil {
		t.Fatal(err)
	}
	if res.StatusCode != 200 {
		t.Fatalf("Expected %d got %d", 200, res.StatusCode)
	}
	// Проверка токена
	body, err := io.ReadAll(res.Body)
	if err != nil {
		t.Fatal(err)
	}
	var resData auth.LoginResponse
	err = json.Unmarshal(body, &resData)
	if err != nil {
		t.Fatal(err)
	}
	if resData.Token == "" {
		t.Fatal("А токен где?")
	}
}

func TestLoginFatal(t *testing.T) {
	r, err := App()
	if err != nil {
		t.Fatalf("Failed to initialize app: %v", err)
	}

	// Создаем тестовый	сервер
	ts := httptest.NewServer(r)
	defer ts.Close()

	data, _ := json.Marshal(&auth.LoginRequest{
		Email:    "a2@a.ru",
		Password: "2",
	})

	res, err := http.Post(ts.URL+"/auth/login", "application/json", bytes.NewReader(data))
	if err != nil {
		t.Fatal(err)
	}
	if res.StatusCode != 401 {
		t.Fatalf("Expected %d got %d", 401, res.StatusCode)
	}
}
