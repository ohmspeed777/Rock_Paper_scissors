package service

import (
	"fmt"
	"time"

	"github.com/dgrijalva/jwt-go"
)

type JWTService interface {
	GenerateToken(email string) string
	ValidateToken(encodetoken string) (*jwt.Token, error)
}

type claimUser struct {
	Id string `json:"id"`
	jwt.StandardClaims
}

type jwtService struct {
	secretKey string
	issue string
}

func NewJwtService() JWTService {
	return &jwtService{
		secretKey: "this-is-secret-key",
		issue: "id",
	}
}

func (jwtService *jwtService) GenerateToken(Id string) string {
	user := &claimUser{
		Id,
		jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Hour * 48).Unix(),
			Issuer:    jwtService.issue,
			IssuedAt:  time.Now().Unix(),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, user)

	tokenEncoded, err := token.SignedString([]byte(jwtService.secretKey))

	if err != nil {
		panic(err)
	}

	return tokenEncoded

}



func (jwtService *jwtService) ValidateToken(encodetoken string) (*jwt.Token, error) {
	return jwt.Parse(encodetoken, func(token *jwt.Token) (interface{}, error) {
		if _, isValid := token.Method.(*jwt.SigningMethodHMAC); !isValid {
			return nil, fmt.Errorf("Invalid token", token.Header["alg"])

		}
		// return the secret key
		return []byte(jwtService.secretKey), nil
	})

}