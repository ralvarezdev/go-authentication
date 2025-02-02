package validator

import (
	"github.com/golang-jwt/jwt/v5"
	gojwttoken "github.com/ralvarezdev/go-jwt/token"
)

// Validator does parsing and validation of JWT tokens
type (
	Validator interface {
		GetToken(rawToken string) (*jwt.Token, error)
		GetClaims(rawToken string) (*jwt.MapClaims, error)
		ValidateClaims(
			rawToken string,
			token gojwttoken.Token,
		) (*jwt.MapClaims, error)
	}
)
