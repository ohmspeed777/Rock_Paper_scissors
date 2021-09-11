package middleware

import (
	"Rock_Paper_Scossors/service"
	"net/http"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"github.com/kamva/mgm/v3"
)

func Protected() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		// header
		const BEARER_SCHEMA = "Bearer "
		authHeader := ctx.GetHeader("Authorization")
		tokenString := authHeader[len(BEARER_SCHEMA):]

		token, err := service.NewJwtService().ValidateToken(tokenString)

		// fmt.Println(token)
		// fmt.Println(token.Valid)

		if !token.Valid || err != nil {
			ctx.JSON(http.StatusUnauthorized, gin.H{
				"msg": "Token is invalid",
				"error": err,
			})
			return 
		}

		claims := token.Claims.(jwt.MapClaims)
		user := service.SingleUser()

		id := claims["id"]
		userErr := mgm.Coll(user).FindByID(id, user)

		if userErr != nil {
			ctx.JSON(http.StatusUnauthorized, gin.H{
				"msg": "This user is not exist.",
				"error": userErr,
			})
			return
		}

		ctx.Set("user", id)
	}
}