package auth_test

import (
	"testing"

	"userDomain/internal/auth"
	"userDomain/internal/user"
)

type MockUserRepository struct{}

func (repo *MockUserRepository) Create(u *user.User) (*user.User, error) {
	return &user.User{
		Email: "a@a.ruy",
	}, nil
}

func (repo *MockUserRepository) FindByEmail(email string) (*user.User, error) {
	return nil, nil
}

func TestRegisterSuccess(t *testing.T) {
	const initialEmail = "a@a.ruy"
	authServiceT := auth.NewAuthService(&MockUserRepository{})
	email, err := authServiceT.Register(initialEmail, "1", "Vasya", "Pechkin")
	if err != nil {
		t.Fatal(err)
	}
	if email != initialEmail {
		t.Fatalf("Email %s do not match %s", email, initialEmail)
	}
}
