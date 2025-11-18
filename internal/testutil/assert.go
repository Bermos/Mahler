package testutil

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

// Helper marks the calling function as a test helper
// and will print a better stack trace if the test fails.
func Helper(t *testing.T) {
	t.Helper()
}

// AssertNoError is a test helper that fails the test if an error is not nil
func AssertNoError(t *testing.T, err error, msgAndArgs ...interface{}) {
	t.Helper()
	require.NoError(t, err, msgAndArgs...)
}

// AssertError is a test helper that fails the test if an error is nil
func AssertError(t *testing.T, err error, msgAndArgs ...interface{}) {
	t.Helper()
	require.Error(t, err, msgAndArgs...)
}

// AssertEqual is a test helper that asserts two values are equal
func AssertEqual(t *testing.T, expected, actual interface{}, msgAndArgs ...interface{}) {
	t.Helper()
	assert.Equal(t, expected, actual, msgAndArgs...)
}

// AssertNotNil is a test helper that asserts a value is not nil
func AssertNotNil(t *testing.T, value interface{}, msgAndArgs ...interface{}) {
	t.Helper()
	assert.NotNil(t, value, msgAndArgs...)
}

// AssertNil is a test helper that asserts a value is nil
func AssertNil(t *testing.T, value interface{}, msgAndArgs ...interface{}) {
	t.Helper()
	assert.Nil(t, value, msgAndArgs...)
}

// AssertLen is a test helper that asserts a slice/map has a specific length
func AssertLen(t *testing.T, object interface{}, length int, msgAndArgs ...interface{}) {
	t.Helper()
	assert.Len(t, object, length, msgAndArgs...)
}

// AssertContains is a test helper that asserts a string contains a substring
func AssertContains(t *testing.T, s, contains string, msgAndArgs ...interface{}) {
	t.Helper()
	assert.Contains(t, s, contains, msgAndArgs...)
}
