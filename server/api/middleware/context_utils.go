package middleware

import (
	"context"
	"errors"
	"github.com/google/uuid"
)

// contextKey is an unexported type to prevent collisions with context keys defined in other packages.
type contextKey string

// UserUUIDKey is the key used to store the authenticated user's UUID in the request context.
// It is exported so handlers in other packages can access it.
const UserUUIDKey contextKey = "userUUID"

func GetUserUUIDFromContext(ctx context.Context)(uuid.UUID, error) {
	val := ctx.Value(UserUUIDKey)
	userID, ok := val.(uuid.UUID)

	if !ok {
		return uuid.Nil, errors.New("user UUID missing or invalid in context")
	}

	return userID, nil
}