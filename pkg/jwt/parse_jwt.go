package jwt

import (
    "errors"

    "github.com/golang-jwt/jwt/v5"
)

// パース＆検証
func ParseJWT(tokenString string,secretKey []byte) (*MyCustomClaims, error) {
    token, err := jwt.ParseWithClaims(tokenString, &MyCustomClaims{}, func(token *jwt.Token) (interface{}, error) {
        return secretKey, nil
    })
    if err != nil || !token.Valid {
        return nil, errors.New("invalid JWT token")
    }
    claims, ok := token.Claims.(*MyCustomClaims)
    if !ok {
        return nil, errors.New("invalid JWT claims")
    }
    return claims, nil
}