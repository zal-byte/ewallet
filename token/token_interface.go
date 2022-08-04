package token

import (
	"ewallet/model"

	"github.com/dgrijalva/jwt-go"
)

type TokenInterface interface {
	Validate(token string) (*jwt.Token, error)
	GenerateToken() *model.Claim
}
