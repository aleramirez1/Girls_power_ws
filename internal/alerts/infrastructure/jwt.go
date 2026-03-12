package infrastructure

import (
	"errors"
	"fmt"

	"github.com/golang-jwt/jwt/v5"
)

func ExtractUserIDFromToken(tokenStr, secret string) (int, error) {
	token, err := jwt.Parse(tokenStr, func(t *jwt.Token) (interface{}, error) {
		if _, ok := t.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("método de firma inesperado: %v", t.Header["alg"])
		}
		return []byte(secret), nil
	})
	if err != nil || !token.Valid {
		return 0, errors.New("token inválido")
	}

	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok {
		return 0, errors.New("claims inválidos")
	}

	rawID, ok := claims["usuario_id"]
	if !ok {
		return 0, errors.New("claim usuario_id no encontrado")
	}

	id, ok := rawID.(float64)
	if !ok {
		return 0, errors.New("usuario_id no es un número")
	}

	return int(id), nil
}