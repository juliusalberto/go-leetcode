package middleware

import (
	"context"
	"fmt"
	"go-leetcode/backend/pkg/response"
	"go-leetcode/backend/internal/authutils"
	"net/http"
	"os"
	"strings"

	"github.com/google/uuid"
)

func AuthMiddleware() func(next http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			authHeader := r.Header.Get("Authorization")
			var userUUID uuid.UUID
			var email string
			// var err error

			// DEV OVERRIDE
			isDevelopment := os.Getenv("GO_ENV") == "development"
			fmt.Println("I'm in the Auth Middleware!")

			if isDevelopment {
				// In development, ALWAYS log the auth header to debug
				if authHeader != "" {
					parts := strings.Split(authHeader, " ")
					if len(parts) == 2 {
						fmt.Println("DEBUG: Received token (first 20 chars):", parts[1][:20])
					}
				}
			}

			if authHeader == "" && isDevelopment {
				fmt.Println("WARN: Dev Override activated - using a pregenerated UUID")
				userUUID = uuid.MustParse("c7413699-cd58-4491-8db5-7de93ba1ac42")
			} else {
				if authHeader == "" {
					response.Error(w, http.StatusUnauthorized, "missing_token", "Authorization header required")
					return
				}

				parts := strings.Split(authHeader, " ")
				if len(parts) != 2 || strings.ToLower(parts[0]) != "bearer" {
					response.Error(w, http.StatusUnauthorized, "invalid_token_format", "Authorization format must be 'Bearer {token}'")
					return
				}

				tokenString := parts[1]

				var err error
				userUUID, email, err = authutils.ParseAndValidateToken(tokenString)

				if err != nil {
					fmt.Printf("JWT validation error: %v\n", err)

					errMsg := "Invalid token"
					errCode := "invalid_token"

					if strings.Contains(err.Error(), "token_expired") {
						errMsg = "Token has expired"
						errCode = "token_expired"
					} else if strings.Contains(err.Error(), "token_signature_invalid") {
						errMsg = "Invalid token signature"
						errCode = "token_signature_invalid"
					}

					response.Error(w, http.StatusUnauthorized, errCode, errMsg)
					return
				}
			}

			if userUUID == uuid.Nil {
				fmt.Println("ERROR: Middleware failed to determine user UUID.")
				response.Error(w, http.StatusInternalServerError, "middleware_error", "Failed to process user identity")
				return
			}

			ctx := context.WithValue(r.Context(), UserUUIDKey, userUUID)
			ctx = context.WithValue(ctx, UserEmailKey, email)
			next.ServeHTTP(w, r.WithContext(ctx))
		})
	}
}