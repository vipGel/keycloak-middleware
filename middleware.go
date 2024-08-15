package keycloak_middleware

import (
	"github.com/golang-jwt/jwt/v4"
	"github.com/labstack/echo/v4"
	"net/http"
	"strings"
)

func EchoClientAuthMiddleware(config *Config) echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			tokenString, err := GetToken(c.Request().Context(), config)
			if err != nil {
				return c.JSON(http.StatusInternalServerError, "Failed to obtain token")
			}

			c.Request().Header.Set("Authorization", "Bearer "+tokenString)

			return next(c)
		}
	}
}

func EchoTokenAuthMiddleware(config *Config) echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			authHeader := c.Request().Header.Get("Authorization")
			if authHeader == "" {
				return c.JSON(http.StatusUnauthorized, "Missing Authorization Header")
			}

			tokenString := strings.TrimPrefix(authHeader, "Bearer ")
			if tokenString == "" {
				return c.JSON(http.StatusUnauthorized, "Invalid Authorization Token")
			}

			token, err := ValidateToken(c.Request().Context(), config, tokenString)
			if err != nil || !token.Valid {
				return c.JSON(http.StatusUnauthorized, "Invalid token")
			}

			claims, ok := token.Claims.(jwt.MapClaims)
			if !ok {
				return c.JSON(http.StatusUnauthorized, "Invalid token")
			}

			c.Set("user", claims)
			return next(c)
		}
	}
}
