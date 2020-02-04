package auth

import (
  "time"

  "github.com/dgrijalva/jwt-go"
)

var (

  // Define a secure key string used
  // as a salt when hashing our tokens.
  // Please make your own way more secure than this,
  // use a randomly generated md5 hash or something.
  key = []byte("mySuperSecretKeyLol")
)

func NewTokenService() *TokenService {
  return &TokenService{}
}

func GetKey() []byte {
  return key
}

// CustomClaims is our custom metadata, which will be hashed
// and sent as the second segment in our JWT
type CustomClaims struct {
  User User
  jwt.StandardClaims
}

// Decode a token string into a token object
func DecodeWithCustomClaims(tokenString string) (*CustomClaims, error) {

  // Parse the token
  token, err := jwt.ParseWithClaims(tokenString, &CustomClaims{}, func(token *jwt.Token) (interface{}, error) {
    return key, nil
  })

  // Validate the token and return the custom claims
  if claims, ok := token.Claims.(*CustomClaims); ok && token.Valid {
    return claims, nil
  } else {
    return nil, err
  }
}
