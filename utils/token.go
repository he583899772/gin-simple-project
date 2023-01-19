package utils

import (
	"fmt"
	"time"

	"github.com/spf13/viper"

	"github.com/dgrijalva/jwt-go"
)

func GenerateAuth() (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"iss": viper.GetString("dependSys.key"),
		"exp": time.Now().Add(time.Hour * 23).Unix(),
	})
	tokenString, err := token.SignedString([]byte(viper.GetString("dependSys.secret")))
	return fmt.Sprintf("Bearer %s", tokenString), err
}
