package jwt

//validar el token

import (
	"errors"
	"strings"

	"github.com/AnthonyCM2K/chatgo/models"
	jwt "github.com/golang-jwt/jwt/v5"
)

var Email string
var IDUsuario string

func ProcesosToken(tk string, JWTSign string) (*models.Claim, bool, string, error) {
	miClave := []byte(JWTSign)
	var claims models.Claim

	splitToken := strings.Split(tk, "Bearer")
	if len(splitToken) != 2 {
		return &claims, false, string(""), errors.New("Formato de token inválido")

	}

	tk = strings.TrimSpace(splitToken[1])

	tkn, err := jwt.ParseWithClaims(tk, &claims, func(token *jwt.Token) (interface{}, error) {
		return miClave, nil
	})
	if err == nil {
		// Rutina que valida contra la base de datos
	}

	if !tkn.Valid {
		return &claims, false, string(""), errors.New("Token Inválido")
	}

	return &claims, false, string(""), err
}
