package testutil

import (
	"errors"
	"testing"
)

// mockT is a mock testing.T that captures calls to Helper() and Errorf()
type mockT struct {
	testing.TB
	helperCalled bool
	errorCalled  bool
	errorMsg     string
}

func (m *mockT) Helper() {
	m.helperCalled = true
}

func (m *mockT) Errorf(format string, args ...interface{}) {
	m.errorCalled = true
	// We don't format the string for simplicity
	m.errorMsg = format
}

func TestAssertEqual(t *testing.T) {
	t.Run("equal_values_pass", func(t *testing.T) {
		AssertEqual(t, 42, 42, "values should be equal")
		AssertEqual(t, "hello", "hello", "strings should be equal")
	})

	t.Run("unequal_values_fail", func(t *testing.T) {
		mock := &mockT{}
		AssertEqual(mock, 42, 43, "test message")
		if !mock.errorCalled {
			t.Error("AssertEqual should call Errorf for unequal values")
		}
		if !mock.helperCalled {
			t.Error("AssertEqual should call Helper")
		}
	})
}

func TestAssertNotEqual(t *testing.T) {
	t.Run("different_values_pass", func(t *testing.T) {
		AssertNotEqual(t, 42, 43, "values should be different")
		AssertNotEqual(t, "hello", "world", "strings should be different")
	})

	t.Run("equal_values_fail", func(t *testing.T) {
		mock := &mockT{}
		AssertNotEqual(mock, 42, 42, "test message")
		if !mock.errorCalled {
			t.Error("AssertNotEqual should call Errorf for equal values")
		}
		if !mock.helperCalled {
			t.Error("AssertNotEqual should call Helper")
		}
	})
}

func TestAssertNil(t *testing.T) {
	t.Run("nil_value_passes", func(t *testing.T) {
		var ptr *int
		AssertNil(t, ptr, "pointer should be nil")
		AssertNil(t, nil, "nil should be nil")
	})

	t.Run("non_nil_value_fails", func(t *testing.T) {
		mock := &mockT{}
		value := 42
		AssertNil(mock, &value, "test message")
		if !mock.errorCalled {
			t.Error("AssertNil should call Errorf for non-nil values")
		}
		if !mock.helperCalled {
			t.Error("AssertNil should call Helper")
		}
	})

	t.Run("non_nil_slice_fails", func(t *testing.T) {
		mock := &mockT{}
		slice := []int{1, 2, 3}
		AssertNil(mock, slice, "test message")
		if !mock.errorCalled {
			t.Error("AssertNil should call Errorf for non-nil slice")
		}
	})
}

func TestAssertNotNil(t *testing.T) {
	t.Run("non_nil_value_passes", func(t *testing.T) {
		value := 42
		AssertNotNil(t, &value, "pointer should not be nil")
	})

	t.Run("nil_value_fails", func(t *testing.T) {
		mock := &mockT{}
		AssertNotNil(mock, nil, "test message")
		if !mock.errorCalled {
			t.Error("AssertNotNil should call Errorf for nil values")
		}
		if !mock.helperCalled {
			t.Error("AssertNotNil should call Helper")
		}
	})

	t.Run("typed_nil_fails", func(t *testing.T) {
		mock := &mockT{}
		var ptr *int
		AssertNotNil(mock, ptr, "test message")
		if !mock.errorCalled {
			t.Error("AssertNotNil should call Errorf for typed nil pointer")
		}
	})
}

func TestAssertNoError(t *testing.T) {
	t.Run("nil_error_passes", func(t *testing.T) {
		AssertNoError(t, nil, "should have no error")
	})

	t.Run("non_nil_error_fails", func(t *testing.T) {
		mock := &mockT{}
		err := errors.New("test error")
		AssertNoError(mock, err, "test message")
		if !mock.errorCalled {
			t.Error("AssertNoError should call Errorf for non-nil error")
		}
		if !mock.helperCalled {
			t.Error("AssertNoError should call Helper")
		}
	})
}

func TestAssertError(t *testing.T) {
	t.Run("non_nil_error_passes", func(t *testing.T) {
		err := errors.New("test error")
		AssertError(t, err, "should have error")
	})

	t.Run("nil_error_fails", func(t *testing.T) {
		mock := &mockT{}
		AssertError(mock, nil, "test message")
		if !mock.errorCalled {
			t.Error("AssertError should call Errorf for nil error")
		}
		if !mock.helperCalled {
			t.Error("AssertError should call Helper")
		}
	})
}

func TestAssertTrue(t *testing.T) {
	t.Run("true_value_passes", func(t *testing.T) {
		AssertTrue(t, true, "value should be true")
	})

	t.Run("false_value_fails", func(t *testing.T) {
		mock := &mockT{}
		AssertTrue(mock, false, "test message")
		if !mock.errorCalled {
			t.Error("AssertTrue should call Errorf for false value")
		}
		if !mock.helperCalled {
			t.Error("AssertTrue should call Helper")
		}
	})
}

func TestAssertFalse(t *testing.T) {
	t.Run("false_value_passes", func(t *testing.T) {
		AssertFalse(t, false, "value should be false")
	})

	t.Run("true_value_fails", func(t *testing.T) {
		mock := &mockT{}
		AssertFalse(mock, true, "test message")
		if !mock.errorCalled {
			t.Error("AssertFalse should call Errorf for true value")
		}
		if !mock.helperCalled {
			t.Error("AssertFalse should call Helper")
		}
	})
}
