package middleware

import (
	"accountability_back/service"
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)

// AuthorizeJWT validates the token from the http request, returning a 401 if it's not valid
func AuthorizeJWT() gin.HandlerFunc {
	return func(c *gin.Context) {
		const BearerSchema = "Bearer "
		authHeader := c.GetHeader("Authorization")
		log.Println("Auth header", authHeader)
		tokenString := ""
		if len(authHeader) > 0 {
			tokenString = authHeader[len(BearerSchema):]
		} else {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
				"message": "No Jwt token present, Please login ",
			})
		}

		token, err := service.JWTAuthService().ValidateToken(tokenString)
		log.Println("err", err)
		log.Println("token", token)
		if token.Valid {
			claims := token.Claims.(jwt.MapClaims)
			log.Println("Claims[Name]: ", claims["name"])
			log.Println("Claims[Email]: ", claims["email"])
			log.Println("Claims[Issuer]: ", claims["iss"])
			log.Println("Claims[IssuedAt]: ", claims["iat"])
			log.Println("Claims[ExpiresAt]: ", claims["exp"])
		} else {
			log.Println(err)
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
				"status": "fail",
				"err":    err.Error(),
			})
		}
	}
}
