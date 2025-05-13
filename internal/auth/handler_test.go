package auth_test

import (
	"bytes"
	"encoding/json"
	"github.com/DATA-DOG/go-sqlmock"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"net/http"
	"net/http/httptest"
	"testing"
	"url/configs"
	"url/internal/auth"
	"url/internal/user"
	"url/pkg/db"
)

func bootstrap() (*auth.AuthHandler, sqlmock.Sqlmock, error) {
	database, mock, err := sqlmock.New()
	if err != nil {
		return nil, nil, err
	}
	GormDb, err := gorm.Open(postgres.New(postgres.Config{
		Conn: database,
	}))
	if err != nil {
		return nil, nil, err
	}
	userRepo := user.NewUserRepository(&db.Db{
		DB: GormDb,
	})
	handler := auth.AuthHandler{
		Config: &configs.Config{
			Auth: configs.AuthConfig{
				Secret: "secret",
			},
		},
		AuthService: auth.NewAuthService(userRepo),
	}
	return &handler, mock, nil
}

func TestLoginSuccess(t *testing.T) {
	handler, mock, err := bootstrap()
	if err != nil {
		t.Fatal(err)
		return
	}
	rows := sqlmock.NewRows([]string{"email", "password"}).AddRow("a4@gmail.com", "$2a$10$7fPVhPjZ9z97WBGY8NGMkeVatz5VfWr9MOFZwcW1e.ssxThfX.0uW")
	mock.ExpectQuery("SELECT").WillReturnRows(rows)
	data, _ := json.Marshal(&auth.LoginRequest{
		Email:    "a4@gmail.com",
		Password: "2",
	})
	reader := bytes.NewReader(data)

	w := httptest.NewRecorder()
	req := httptest.NewRequest(http.MethodPost, "/auth/login", reader)
	handler.Login()(w, req)
	if w.Code != http.StatusOK {
		t.Errorf("Login failed. Expected %d. Got %d", http.StatusOK, w.Code)
		return
	}
}
func TestRegSuccess(t *testing.T) {
	handler, mock, err := bootstrap()
	if err != nil {
		t.Fatal(err)
		return
	}
	rows := sqlmock.NewRows([]string{"email", "password", "name"})
	mock.ExpectQuery("SELECT").WillReturnRows(rows)
	mock.ExpectBegin()
	mock.ExpectQuery("INSERT").WillReturnRows(sqlmock.NewRows([]string{"id"}).AddRow(1))
	mock.ExpectCommit()
	data, _ := json.Marshal(&auth.RegisterRequest{
		Email:    "a4@gmail.com",
		Password: "2",
		Name:     "Толя",
	})
	reader := bytes.NewReader(data)

	w := httptest.NewRecorder()
	req := httptest.NewRequest(http.MethodPost, "/auth/register", reader)
	handler.Register()(w, req)
	if w.Code != http.StatusOK {
		t.Errorf("Register failed. Expected %d. Got %d", http.StatusOK, w.Code)
		return
	}
}
