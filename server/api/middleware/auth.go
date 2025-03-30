package middleware

import (
	"context"
	"fmt"
	"go-leetcode/backend/pkg/response"
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
			// var err error

			// DEV OVERRIDE
			isDevelopment := os.Getenv("GO_ENV") == "development"
			fmt.Println("I'm in the Auth Middleware!")

			if authHeader == "" && isDevelopment {
				fmt.Println("WARN: Dev Override activated - using a pregenerated UUID")
				userUUID = uuid.MustParse("c7413699-cd58-4491-8db5-7de93ba1ac42")
			} else {
				if authHeader == "" {
					response.Error(w, http.StatusUnauthorized, "missing_token", "Authorization header required")
					return
				}

				parts := strings.Split(authHeader, " ")
				if len(parts) != 2 || strings.ToLower(parts[0]) != "Bearer" {
					response.Error(w, http.StatusUnauthorized, "invalid_token_format", "Authorization format must be 'Bearer {token}'")
					return
				}

				// !! TODO: Implement REAL JWT Validation Here for production !!
			}

			if userUUID == uuid.Nil {
				fmt.Println("ERROR: Middleware failed to determine user UUID.")
				response.Error(w, http.StatusInternalServerError, "middleware_error", "Failed to process user identity")
				return
			}

			ctx := context.WithValue(r.Context(), UserUUIDKey, userUUID)
			next.ServeHTTP(w, r.WithContext(ctx))
		})
	}
}