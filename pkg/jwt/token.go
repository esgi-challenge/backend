package jwt

import (
	"encoding/json"
	"errors"
	"time"

	"github.com/esgi-challenge/backend/internal/models"
	j "github.com/golang-jwt/jwt/v5"
)

type UserClaim struct {
	Exp  float64
	User models.User
}

func Generate(secretKey string, user *models.User) (string, error) {
	user.Password = ""

	token := j.NewWithClaims(
		j.SigningMethodHS256,
		j.MapClaims{
			"user": user,
			"exp":  time.Now().Add(time.Hour * 24).Unix(),
		},
	)

	tokenString, err := token.SignedString([]byte(secretKey))

	if err != nil {
		return "", err
	}

	return tokenString, nil
}

func DecryptToken(secretKey string, tokenString string) (*models.User, error) {
	token, err := j.Parse(tokenString, func(token *j.Token) (interface{}, error) {
		return []byte(secretKey), nil
	})

	if err != nil {
		return nil, err
	}

	if !token.Valid {
		return nil, errors.New("jwt is not valid")
	}

	if claims, ok := token.Claims.(j.MapClaims); ok {
		dict, err := json.Marshal(claims)

		if err != nil {
			return nil, err
		}

		userClaim := &UserClaim{}

		if err := json.Unmarshal(dict, &userClaim); err != nil {
			return nil, err
		}

		return &userClaim.User, nil
	}

	return nil, nil

}
