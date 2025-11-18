package testutil

import (
	"testing"
	"time"
)

func TestNewTestContext(t *testing.T) {
	t.Run("creates_valid_context", func(t *testing.T) {
		ctx := NewTestContext(t)
		AssertNotNil(t, ctx, "context should not be nil")

		// Verify context has a deadline
		_, hasDeadline := ctx.Deadline()
		AssertTrue(t, hasDeadline, "context should have a deadline")
	})

	t.Run("context_has_proper_timeout", func(t *testing.T) {
		ctx := NewTestContext(t)
		deadline, ok := ctx.Deadline()
		AssertTrue(t, ok, "context should have deadline")

		// Deadline should be in the future
		AssertTrue(t, time.Until(deadline) > 0, "deadline should be in the future")
	})
}

func TestNewTestContextWithCancel(t *testing.T) {
	t.Run("creates_valid_context_with_cancel", func(t *testing.T) {
		ctx, cancel := NewTestContextWithCancel(t)
		defer cancel()

		AssertNotNil(t, ctx, "context should not be nil")
		AssertNotNil(t, cancel, "cancel function should not be nil")

		// Verify context has a deadline
		_, hasDeadline := ctx.Deadline()
		AssertTrue(t, hasDeadline, "context should have a deadline")
	})

	t.Run("cancel_function_works", func(t *testing.T) {
		ctx, cancel := NewTestContextWithCancel(t)

		// Cancel the context
		cancel()

		// Wait a bit to ensure cancellation propagates
		time.Sleep(10 * time.Millisecond)

		// Context should be done
		select {
		case <-ctx.Done():
			// Success - context was cancelled
		default:
			t.Error("context should be cancelled")
		}
	})
}
