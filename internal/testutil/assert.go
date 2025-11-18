package testutil

import (
	"reflect"
	"testing"
)

// AssertEqual checks if two values are equal
func AssertEqual[T comparable](t testing.TB, got, want T, message string) {
	t.Helper()
	if got != want {
		t.Errorf("%s: got %v, want %v", message, got, want)
	}
}

// AssertNotEqual checks if two values are not equal
func AssertNotEqual[T comparable](t testing.TB, got, notWant T, message string) {
	t.Helper()
	if got == notWant {
		t.Errorf("%s: got %v, expected it to be different", message, got)
	}
}

// AssertNil checks if a value is nil
func AssertNil(t testing.TB, got interface{}, message string) {
	t.Helper()
	if got == nil {
		return
	}
	// Use reflection to check if the underlying value is nil
	// This handles typed nil pointers correctly
	v := reflect.ValueOf(got)
	if v.Kind() == reflect.Ptr || v.Kind() == reflect.Interface ||
		v.Kind() == reflect.Slice || v.Kind() == reflect.Map ||
		v.Kind() == reflect.Chan || v.Kind() == reflect.Func {
		if !v.IsNil() {
			t.Errorf("%s: got %v, want nil", message, got)
		}
	} else {
		t.Errorf("%s: got %v, want nil", message, got)
	}
}

// AssertNotNil checks if a value is not nil
func AssertNotNil(t testing.TB, got interface{}, message string) {
	t.Helper()
	if got == nil {
		t.Errorf("%s: got nil, want non-nil value", message)
		return
	}
	// Use reflection to check if the underlying value is nil
	v := reflect.ValueOf(got)
	if v.Kind() == reflect.Ptr || v.Kind() == reflect.Interface ||
		v.Kind() == reflect.Slice || v.Kind() == reflect.Map ||
		v.Kind() == reflect.Chan || v.Kind() == reflect.Func {
		if v.IsNil() {
			t.Errorf("%s: got nil, want non-nil value", message)
		}
	}
}

// AssertNoError checks if an error is nil
func AssertNoError(t testing.TB, err error, message string) {
	t.Helper()
	if err != nil {
		t.Errorf("%s: unexpected error: %v", message, err)
	}
}

// AssertError checks if an error is not nil
func AssertError(t testing.TB, err error, message string) {
	t.Helper()
	if err == nil {
		t.Errorf("%s: expected error, got nil", message)
	}
}

// AssertTrue checks if a boolean is true
func AssertTrue(t testing.TB, got bool, message string) {
	t.Helper()
	if !got {
		t.Errorf("%s: got false, want true", message)
	}
}

// AssertFalse checks if a boolean is false
func AssertFalse(t testing.TB, got bool, message string) {
	t.Helper()
	if got {
		t.Errorf("%s: got true, want false", message)
	}
}
