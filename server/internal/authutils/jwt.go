package authutils

import (
	"errors"
	"fmt"
	"os"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
)


type CustomClaims struct {
	UserID 		string `json:"sub"`
	Email		string `json:"email"`
	jwt.RegisteredClaims
}

func getJwtSecret() ([]byte, error) {
	secret := os.Getenv("SUPABASE_JWT_SECRET")
	if secret == "" {
		return nil, errors.New("SUPABASE_JWT_SECRET environment variable not set or empty")
	}
	return []byte(secret), nil
}

func Initialize() {
	isDevelopment := os.Getenv("GO_ENV") == "development"
	if !isDevelopment {
		_, err := getJwtSecret()
		if err != nil {
			panic("FATAL: SUPABASE_JWT_SECRET env not set or empty")
		}
	}
}

func ParseAndValidateToken(tokenString string) (uuid.UUID, string, error) {
	jwtSecret, err := getJwtSecret()
	if err != nil {
		return uuid.Nil, "", fmt.Errorf("config_error: %w", err)
	}

	claims := &CustomClaims{}

	token, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token)(interface{}, error) {

		// ensure only HMAC token is accepted
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}

		return jwtSecret, nil
	}, jwt.WithAudience("authenticated"), jwt.WithLeeway(5*time.Second))

	if err != nil {
		if errors.Is(err, jwt.ErrTokenExpired) {
			return uuid.Nil, "", fmt.Errorf("token_expired %w", err)
		} else if errors.Is(err, jwt.ErrTokenNotValidYet) {
			return uuid.Nil, "", fmt.Errorf("token_not_active: %w", err)
		} else if errors.Is(err, jwt.ErrTokenSignatureInvalid) {
			return uuid.Nil, "", fmt.Errorf("token_signature_invalid: %w", err)
		} 

		return uuid.Nil, "", fmt.Errorf("invalid_token: %w", err)
	}

	if !token.Valid {
		return uuid.Nil, "", errors.New("invalid_token")
	}

	if claims.UserID == "" {
		return uuid.Nil, "", errors.New("missing_sub_claim")
	}

	userID, uuidErr := uuid.Parse(claims.UserID)
	if uuidErr != nil {
		return uuid.Nil, "", fmt.Errorf("invalid_sub_claim_format: %w", uuidErr)
	}

	return userID, claims.Email, nil
}