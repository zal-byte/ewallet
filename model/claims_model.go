package model

import "github.com/dgrijalva/jwt-go"

type Claim struct {
	jwt.StandardClaims
	Username string `json:"username"`
	Token    string `json:"generated_token"`
}
