package utils

import (
	"errors"
	"os"
	"time"
	"github.com/golang-jwt/jwt/v5"
)

var secretKey = []byte(os.Getenv("secret_key"))


// struktur data dari jwt
type CustomClaims struct{
	Username string `json:"username"`
	jwt.RegisteredClaims
}



// fungsi untuk generate jwt
func generateJWT(username string)(string, error){
	claims := CustomClaims{
		Username: username,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add( 1 * time.Hour)),
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(secretKey)
}

// fungsi untuk validasi token jwt
func validateJwt(tokenString string) (*CustomClaims, error){
	token, err := jwt.ParseWithClaims(tokenString, &CustomClaims{}, func(t *jwt.Token) (interface{}, error) {
		return secretKey, nil
	})
	if err != nil{
		return nil, err
	}
	
	claims, ok := token.Claims.(*CustomClaims)
	if !ok || !token.Valid {
		return nil, errors.New("token tidak valid")
	}

	return claims, nil
}