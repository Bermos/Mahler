package testutil

import (
	"context"
	"testing"
	"time"
)

// NewTestContext creates a test context with a timeout
func NewTestContext(t *testing.T) context.Context {
	t.Helper()
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	t.Cleanup(cancel)
	return ctx
}

// NewTestContextWithCancel creates a test context with a cancel function
func NewTestContextWithCancel(t *testing.T) (context.Context, context.CancelFunc) {
	t.Helper()
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	t.Cleanup(cancel)
	return ctx, cancel
}
