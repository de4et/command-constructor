package api

import (
	"fmt"
	"os"

	"github.com/de4et/command-constructor/types"
	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v5"
)

func JWTAuth(c *fiber.Ctx) error {
	fmt.Println("-- JWT authenticating")

	token, ok := c.GetReqHeaders()["X-Api-Token"]
	if !ok {
		return fmt.Errorf("unauthorized")
	}
	fmt.Println(token)

	claims, err := parseToken(token[0])
	if err != nil {
		fmt.Println("Failed to parse JWT token:", err)
		return fmt.Errorf("unauthorized")
	}
	fmt.Println(claims)
	return c.JSON("im here") // FIXME

}

func parseToken(tokenStr string) (jwt.MapClaims, error) { // lower case
	token, err := jwt.Parse(tokenStr, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}

		secret := os.Getenv("JWT_SECRET")
		fmt.Println("NEVER PRINT OUT SECRET", secret)
		return []byte(secret), nil
	})
	if err != nil {
		return nil, fmt.Errorf("unauthorized")
	}

	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok {
		return nil, fmt.Errorf("unauthorized")
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
