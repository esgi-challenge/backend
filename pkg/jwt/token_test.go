package jwt

import (
	"testing"
	"time"

	"github.com/esgi-challenge/backend/internal/models"
	j "github.com/golang-jwt/jwt/v5"
	"github.com/stretchr/testify/assert"
)

func TestGenerate(t *testing.T) {
	secretKey := "secret"
	user := &models.User{
		Firstname: "John",
		Lastname:  "Doe",
		Email:     "john.doe@example.com",
		UserKind:  new(models.UserKind),
	}
	*user.UserKind = models.ADMINISTRATOR

	token, err := Generate(secretKey, user)
	assert.NoError(t, err)
	assert.NotEmpty(t, token)
}

func TestDecryptToken(t *testing.T) {
	secretKey := "secret"
	user := &models.User{
		Firstname: "John",
		Lastname:  "Doe",
		Email:     "john.doe@example.com",
		UserKind:  new(models.UserKind),
	}
	*user.UserKind = models.ADMINISTRATOR

	token, err := Generate(secretKey, user)
	assert.NoError(t, err)
	assert.NotEmpty(t, token)

	decryptedUser, err := DecryptToken(secretKey, token)
	assert.NoError(t, err)
	assert.NotNil(t, decryptedUser)
	assert.Equal(t, user.Firstname, decryptedUser.Firstname)
	assert.Equal(t, user.Lastname, decryptedUser.Lastname)
	assert.Equal(t, user.Email, decryptedUser.Email)
	assert.Equal(t, *user.UserKind, *decryptedUser.UserKind)
}

func TestDecryptToken_InvalidToken(t *testing.T) {
	secretKey := "secret"
	invalidToken := "invalidToken"

	decryptedUser, err := DecryptToken(secretKey, invalidToken)
	assert.Error(t, err)
	assert.Nil(t, decryptedUser)
}

func TestDecryptToken_ExpiredToken(t *testing.T) {
	secretKey := "secret"
	user := &models.User{
		Firstname: "John",
		Lastname:  "Doe",
		Email:     "john.doe@example.com",
		UserKind:  new(models.UserKind),
	}
	*user.UserKind = models.ADMINISTRATOR

	token := j.NewWithClaims(
		j.SigningMethodHS256,
		j.MapClaims{
			"user": user,
			"exp":  time.Now().Add(time.Second * 1).Unix(),
		},
	)
	tokenString, err := token.SignedString([]byte(secretKey))
	assert.NoError(t, err)
	assert.NotEmpty(t, tokenString)

	time.Sleep(time.Second * 2)

	decryptedUser, err := DecryptToken(secretKey, tokenString)
	assert.Error(t, err)
	assert.Nil(t, decryptedUser)
}

func TestDecryptToken_TamperedToken(t *testing.T) {
	secretKey := "secret"
	user := &models.User{
		Firstname: "John",
		Lastname:  "Doe",
		Email:     "john.doe@example.com",
		UserKind:  new(models.UserKind),
	}
	*user.UserKind = models.ADMINISTRATOR

	token, err := Generate(secretKey, user)
	assert.NoError(t, err)
	assert.NotEmpty(t, token)

	tamperedToken := token + "tampered"

	decryptedUser, err := DecryptToken(secretKey, tamperedToken)
	assert.Error(t, err)
	assert.Nil(t, decryptedUser)
}
