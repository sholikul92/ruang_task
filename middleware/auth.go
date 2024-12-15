package middleware

import (
	"a21hc3NpZ25tZW50/model"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"
)

func Auth() gin.HandlerFunc {
	return func(c *gin.Context) {
		cookie, err := c.Cookie("session_token")
		if err != nil {
			if err == http.ErrNoCookie {
				fmt.Println("error no cookie: ", err)
				c.AbortWithStatusJSON(http.StatusUnauthorized, model.CreateHandlerResponseError(err.Error()))
				return
			}

			fmt.Println("error request cookie : ", err)
			c.AbortWithStatusJSON(http.StatusBadRequest, model.CreateHandlerResponseError(err.Error()))
			return
		}

		var claims model.Claims
		jwtToken, err := jwt.ParseWithClaims(cookie, &claims, func(t *jwt.Token) (interface{}, error) {
			return model.JwtKey, nil
		})

		if err != nil {
			if err == jwt.ErrSignatureInvalid {
				c.AbortWithStatusJSON(http.StatusUnauthorized, model.CreateHandlerResponseError(err.Error()))
				return
			}

			c.AbortWithStatusJSON(http.StatusBadRequest, model.CreateHandlerResponseError(err.Error()))
			fmt.Println(jwtToken)
			return
		}

		if !jwtToken.Valid {
			c.AbortWithStatus(http.StatusUnauthorized)
			return
		}

		c.Next()
	}
}
