package handlers

import (
	"errors"
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"log/slog"
)

const BEARER_SCHEMA = "Bearer "

func VerifyJWT() gin.HandlerFunc {

	secretKey := "secret"
	return func(ctx *gin.Context) {
		jwtToken := ctx.GetHeader("Authorization")
		tokenStr := jwtToken[len(BEARER_SCHEMA):]
		_, err := jwt.Parse(tokenStr, func(token *jwt.Token) (interface{}, error) {
			if _, isvalid := token.Method.(*jwt.SigningMethodHMAC); !isvalid {
				slog.Error("Invalid token")
				return nil, errors.New("Invalid token")
			}
			return []byte(secretKey), nil
		})

		if err != nil {
			slog.Error("Verify token fail", slog.String("Err", err.Error()))
			ctx.AbortWithStatus(401)
		}
		ctx.Next()
	}

}
