package context

import (
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	gojwt "github.com/ralvarezdev/go-jwt"
)

// SetCtxTokenClaims sets the token claims in the context
func SetCtxTokenClaims(ctx *gin.Context, claims *jwt.MapClaims) {
	ctx.Set(gojwt.CtxTokenClaimsKey, *claims)
}

// GetCtxTokenClaims tries to get the token claims from the context
func GetCtxTokenClaims(ctx *gin.Context) (*jwt.MapClaims, error) {
	// Get the token claims from the context
	value := ctx.Value(gojwt.CtxTokenClaimsKey)
	if value == nil {
		return nil, gojwt.ErrMissingTokenClaimsInContext
	}

	// Check the type of the value
	claims, ok := value.(jwt.MapClaims)
	if !ok {
		return nil, gojwt.ErrUnexpectedTokenClaimsTypeInContext
	}

	return &claims, nil
}
