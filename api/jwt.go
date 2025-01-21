package api

import (
	"fmt"
	"log"
	"os"

	"github.com/de4et/command-constructor/db"
	"github.com/de4et/command-constructor/types"
	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v5"
)

func JWTAuth(store *db.Store) func(c *fiber.Ctx) error {
	return func(c *fiber.Ctx) error {
		token := c.Cookies("apiToken")
		if token == "" {
			return ErrUnauthorized()
		}

		claims, err := parseToken(token)
		if err != nil {
			fmt.Println("Failed to parse JWT token:", err)
			c.ClearCookie("apiToken")
			return ErrUnauthorized()
		}

		user, err := store.User.GetUserByID(c.Context(), claims["id"].(string))
		if err != nil {
			log.Println(err)
			return ErrUnauthorized()
		}

		c.Context().SetUserValue("user", user)
		return c.Next()
	}
}

func parseToken(tokenStr string) (jwt.MapClaims, error) { // lower case
	token, err := jwt.Parse(tokenStr, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}

		secret := os.Getenv("JWT_SECRET")
		return []byte(secret), nil
	})
	if err != nil {
		return nil, ErrUnauthorized()
	}

	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok {
		return nil, ErrUnauthorized()
	}
	return claims, nil
}

func makeTokenFromUser(user *types.User) string {
	claims := jwt.MapClaims{
		"id":    user.ID,
		"email": user.Email,
		"name":  user.Name,
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	secret := os.Getenv("JWT_SECRET")

	tokenStr, err := token.SignedString([]byte(secret))
	if err != nil {
		fmt.Println("failed to sign token with secret")
	}
	return tokenStr
}
func AuthMiddleware(store *db.Store) func(c *fiber.Ctx) error {
	return func(c *fiber.Ctx) error {
		var user *types.User = nil
		c.Context().SetUserValue("user", user)

		token := c.Cookies("apiToken")
		if token == "" {
			return c.Next()
		}

		claims, err := parseToken(token)
		if err != nil {
			fmt.Println("Failed to parse JWT token:", err)
			c.ClearCookie("apiToken")
			return c.Next()
		}

		user, err = store.User.GetUserByID(c.Context(), claims["id"].(string))
		if err != nil {
			log.Println(err)
			return c.Next()
		}

		c.Context().SetUserValue("user", user)
		return c.Next()
	}
}
