package jwt_test

import (
	"testing"
	"url/pkg/jwt"
)

func TestJWT_Create(t *testing.T) {
	const email = "a@a.ru"
	JWTService := jwt.NewJWT("/2+XnmJGz1j3ehIVI/5P9k1+CghrE3DcS7rnT+qar5w=")
	token, err := JWTService.Create(jwt.JWTData{
		Email: email,
	})
	if err != nil {
		t.Fatal(err)
	}
	isValid, data := JWTService.Parse(token)
	if !isValid {
		t.Fatal("Invalid token")
	}
	if data.Email != email {
		t.Fatalf("Email %s not equal to %s ", data.Email, email)
	}
}
