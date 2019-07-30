package jwtauth

import (
	"errors"
	"fmt"
	"time"

	jwt "github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
)

var (
	// ErrMissingHeader means the `Authorization` header was empty.
	ErrMissingHeader = errors.New("The length of the `Authorization` header is zero")
)

// Claims is the context of the JSON web token.
type Claims struct {
	UserID   string `json:"user_id"`
	Username string `json:"username"`
	jwt.StandardClaims
}

// secretFunc validates the secret format.
func secretFunc(secret string) jwt.Keyfunc {
	return func(token *jwt.Token) (interface{}, error) {
		// Make sure the `alg` is what we except.
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, jwt.ErrSignatureInvalid
		}

		return []byte(secret), nil
	}
}

// ParseToken validates the token with the specified secret,
// and returns the context if the token was valid.
func ParseToken(tokenString string, secret string) (*Claims, error) {

	// Parse the token.
	// token, err := jwt.Parse(tokenString, secretFunc(secret))
	tokenClaims, err := jwt.ParseWithClaims(tokenString, &Claims{}, secretFunc(secret))
	if tokenClaims != nil {
		if claims, ok := tokenClaims.Claims.(*Claims); ok && tokenClaims.Valid {
			return claims, nil
		}
	}

	return nil, err

}

// ParseRequest gets the token from the header and
// pass it to the Parse function to parses the token.
func ParseRequest(c *gin.Context) (*Claims, error) {
	header := c.Request.Header.Get("Authorization")

	// Load the jwt secret from config
	secret := viper.GetString("jwt.secret")

	if len(header) == 0 {
		return &Claims{}, ErrMissingHeader
	}

	var t string
	// Parse the header to get the token part.
	fmt.Sscanf(header, "Bearer %s", &t)
	return ParseToken(t, secret)
}

// GenerateToken signs the context with the specified secret.
func GenerateToken(ctx *gin.Context, c Claims, secret string) (tokenString string, err error) {

	// Load the jwt secret from the Gin config if the secret isn't specified.
	if secret == "" {
		secret = viper.GetString("jwt.secret")
	}

	gapTime := viper.GetDuration("jwt.exp_gap_time")
	if gapTime == 0 {
		gapTime = 24
	}
	nowTime := time.Now()
	expireTime := nowTime.Add(gapTime * time.Hour) // 设置token gapTime小时后过期

	claims := Claims{
		c.UserID,
		c.Username,
		jwt.StandardClaims{
			ExpiresAt: expireTime.Unix(),
			Issuer:    "entere-micro-blog",
		},
	}

	// The token content.
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	// Sign the token with the specified secret.
	tokenString, err = token.SignedString([]byte(secret))

	return
}
