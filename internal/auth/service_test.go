package auth_test

import (
	"testing"
	"url/internal/auth"
	"url/internal/user"
)

type MockUserRepository struct {
}

func (repo *MockUserRepository) Create(*user.User) (*user.User, error) {
	return &user.User{
		Email: "a@a.ru",
	}, nil

}

func (repo *MockUserRepository) FindUserByEmail(email string) (*user.User, error) {
	return nil, nil

}

func TestRegisterSuccess(t *testing.T) {
	const email = "a@a.ru"
	authService := auth.NewAuthService(&MockUserRepository{})

	emailReg, err := authService.Register(email, "1", "Паша")
	if err != nil {
		t.Fatal(err)
	}
	if emailReg != email {
		t.Fatalf("email register failed, expected %s, got %s", email, emailReg)
	}
}
