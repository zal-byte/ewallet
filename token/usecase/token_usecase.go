package usecase

import (
	"ewallet/model"
	"fmt"
	"os"
	"time"

	"github.com/dgrijalva/jwt-go"
)

var SECRET_KEY = ""

type TokenUsecase struct {
}

func CreateTokenUsecase() TokenUsecase {
	if os.Getenv("SECRET_KEY") == "" {
		SECRET_KEY = "secret"
	} else {
		SECRET_KEY = os.Getenv("SECRET_KEY")
	}

	return TokenUsecase{}
}

func (e *TokenUsecase) GenerateToken() *model.Claim {
	expirationTime := time.Now().Add(5 * time.Minute)

	claims := model.Claim{
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expirationTime.Unix(),
		},
	}
	
	sign := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	token, err := sign.SignedString([]byte(SECRET_KEY))
	claims.Token = token

	if err != nil {
		panic(err)
	}
	return &claims

}

func Validate(token string) (*jwt.Token, error) {
	return jwt.Parse(token, func(t *jwt.Token) (interface{}, error) {
		if _, isvalid := t.Method.(*jwt.SigningMethodHMAC); !isvalid {
			return nil, fmt.Errorf("invalid token %s", t.Header["alg"])
		}
		return []byte(SECRET_KEY), nil
	})
}
