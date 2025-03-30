package middleware

import (
	"context"
	"errors"
	"fmt"

	"github.com/google/uuid"
)

// contextKey is an unexported type to prevent collisions with context keys defined in other packages.
type contextKey string

// UserUUIDKey is the key used to store the authenticated user's UUID in the request context.
// It is exported so handlers in other packages can access it.
const UserUUIDKey contextKey = "userUUID"
const UserEmailKey contextKey = "userEmail"

func GetUserUUIDFromContext(ctx context.Context)(uuid.UUID, error) {
	val := ctx.Value(UserUUIDKey)
	userID, ok := val.(uuid.UUID)
	fmt.Printf("GetUserUUIDFromContext: Value for key %v is: %v (Type: %T)\n", UserUUIDKey, val, val)

	if !ok {
		return uuid.Nil, errors.New("user UUID missing or invalid in context")
	}

	return userID, nil
}

func GetStringFromContext(ctx context.Context, key contextKey)(string, error) {
	val := ctx.Value(key)
	stringVal, ok := val.(string)
	fmt.Printf("GetStringFromContext: Value for key %v is: %v (Type: %T)\n", key, val, val)

	if !ok {
		return "", errors.New("string missing or invalid in context")
	}

	return stringVal, nil
}