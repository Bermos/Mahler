package config

import (
	"os"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestLoad(t *testing.T) {
	t.Parallel()

	// Clear environment
	cleanup := setupTestEnv(t, nil)
	defer cleanup()

	cfg, err := Load()
	require.NoError(t, err)
	require.NotNil(t, cfg)

	// Check defaults
	assert.Equal(t, "", cfg.Server.Host)
	assert.Equal(t, 8080, cfg.Server.Port)
	assert.Equal(t, "sqlite", cfg.Database.Driver)
	assert.Equal(t, "mahler.db", cfg.Database.ConnectionString)
	assert.Equal(t, "info", cfg.Logging.Level)
	assert.Equal(t, "json", cfg.Logging.Format)
}

func TestLoad_WithEnvironmentVariables(t *testing.T) {
	t.Parallel()

	env := map[string]string{
		"SERVER_HOST":              "localhost",
		"SERVER_PORT":              "9090",
		"DB_DRIVER":                "postgres",
		"DB_CONNECTION_STRING":     "postgres://localhost/test",
		"DB_MAX_OPEN_CONNS":        "50",
		"DB_MAX_IDLE_CONNS":        "10",
		"LOG_LEVEL":                "debug",
		"LOG_FORMAT":               "text",
	}

	cleanup := setupTestEnv(t, env)
	defer cleanup()

	cfg, err := Load()
	require.NoError(t, err)

	assert.Equal(t, "localhost", cfg.Server.Host)
	assert.Equal(t, 9090, cfg.Server.Port)
	assert.Equal(t, "postgres", cfg.Database.Driver)
	assert.Equal(t, "postgres://localhost/test", cfg.Database.ConnectionString)
	assert.Equal(t, 50, cfg.Database.MaxOpenConns)
	assert.Equal(t, 10, cfg.Database.MaxIdleConns)
	assert.Equal(t, "debug", cfg.Logging.Level)
	assert.Equal(t, "text", cfg.Logging.Format)
}

func TestServerConfig_Validate(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name        string
		config      ServerConfig
		expectError bool
		errorMsg    string
	}{
		{
			name: "valid config",
			config: ServerConfig{
				Host:         "localhost",
				Port:         8080,
				ReadTimeout:  15 * time.Second,
				WriteTimeout: 15 * time.Second,
				IdleTimeout:  60 * time.Second,
			},
			expectError: false,
		},
		{
			name: "invalid port - too low",
			config: ServerConfig{
				Port: 0,
			},
			expectError: true,
			errorMsg:    "invalid port",
		},
		{
			name: "invalid port - too high",
			config: ServerConfig{
				Port: 70000,
			},
			expectError: true,
			errorMsg:    "invalid port",
		},
		{
			name: "negative read timeout",
			config: ServerConfig{
				Port:        8080,
				ReadTimeout: -1 * time.Second,
			},
			expectError: true,
			errorMsg:    "read timeout cannot be negative",
		},
		{
			name: "negative write timeout",
			config: ServerConfig{
				Port:         8080,
				WriteTimeout: -1 * time.Second,
			},
			expectError: true,
			errorMsg:    "write timeout cannot be negative",
		},
		{
			name: "negative idle timeout",
			config: ServerConfig{
				Port:        8080,
				IdleTimeout: -1 * time.Second,
			},
			expectError: true,
			errorMsg:    "idle timeout cannot be negative",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			err := tt.config.Validate()

			if tt.expectError {
				require.Error(t, err)
				assert.Contains(t, err.Error(), tt.errorMsg)
			} else {
				require.NoError(t, err)
			}
		})
	}
}

func TestDatabaseConfig_Validate(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name        string
		config      DatabaseConfig
		expectError bool
		errorMsg    string
	}{
		{
			name: "valid sqlite config",
			config: DatabaseConfig{
				Driver:           "sqlite",
				ConnectionString: "test.db",
				MaxOpenConns:     25,
				MaxIdleConns:     5,
			},
			expectError: false,
		},
		{
			name: "valid postgres config",
			config: DatabaseConfig{
				Driver:           "postgres",
				ConnectionString: "postgres://localhost/test",
				MaxOpenConns:     50,
				MaxIdleConns:     10,
			},
			expectError: false,
		},
		{
			name: "invalid driver",
			config: DatabaseConfig{
				Driver:           "mongodb",
				ConnectionString: "mongodb://localhost",
				MaxOpenConns:     10,
			},
			expectError: true,
			errorMsg:    "invalid database driver",
		},
		{
			name: "empty connection string",
			config: DatabaseConfig{
				Driver:       "sqlite",
				MaxOpenConns: 10,
			},
			expectError: true,
			errorMsg:    "connection string is required",
		},
		{
			name: "invalid max open connections",
			config: DatabaseConfig{
				Driver:           "sqlite",
				ConnectionString: "test.db",
				MaxOpenConns:     0,
			},
			expectError: true,
			errorMsg:    "max open connections must be at least 1",
		},
		{
			name: "negative max idle connections",
			config: DatabaseConfig{
				Driver:           "sqlite",
				ConnectionString: "test.db",
				MaxOpenConns:     10,
				MaxIdleConns:     -1,
			},
			expectError: true,
			errorMsg:    "max idle connections cannot be negative",
		},
		{
			name: "max idle exceeds max open",
			config: DatabaseConfig{
				Driver:           "sqlite",
				ConnectionString: "test.db",
				MaxOpenConns:     10,
				MaxIdleConns:     20,
			},
			expectError: true,
			errorMsg:    "max idle connections cannot exceed max open connections",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			err := tt.config.Validate()

			if tt.expectError {
				require.Error(t, err)
				assert.Contains(t, err.Error(), tt.errorMsg)
			} else {
				require.NoError(t, err)
			}
		})
	}
}

func TestLoggingConfig_Validate(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name        string
		config      LoggingConfig
		expectError bool
		errorMsg    string
	}{
		{
			name: "valid config - info/json",
			config: LoggingConfig{
				Level:  "info",
				Format: "json",
			},
			expectError: false,
		},
		{
			name: "valid config - debug/text",
			config: LoggingConfig{
				Level:  "debug",
				Format: "text",
			},
			expectError: false,
		},
		{
			name: "invalid log level",
			config: LoggingConfig{
				Level:  "trace",
				Format: "json",
			},
			expectError: true,
			errorMsg:    "invalid log level",
		},
		{
			name: "invalid log format",
			config: LoggingConfig{
				Level:  "info",
				Format: "xml",
			},
			expectError: true,
			errorMsg:    "invalid log format",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			err := tt.config.Validate()

			if tt.expectError {
				require.Error(t, err)
				assert.Contains(t, err.Error(), tt.errorMsg)
			} else {
				require.NoError(t, err)
			}
		})
	}
}

func TestGetEnv(t *testing.T) {
	t.Parallel()

	t.Run("returns default when env not set", func(t *testing.T) {
		value := getEnv("NON_EXISTENT_VAR", "default")
		assert.Equal(t, "default", value)
	})

	t.Run("returns env value when set", func(t *testing.T) {
		key := "TEST_VAR_" + t.Name()
		expected := "test_value"
		os.Setenv(key, expected)
		defer os.Unsetenv(key)

		value := getEnv(key, "default")
		assert.Equal(t, expected, value)
	})
}

func TestGetEnvAsInt(t *testing.T) {
	t.Parallel()

	t.Run("returns default when env not set", func(t *testing.T) {
		value := getEnvAsInt("NON_EXISTENT_INT", 42)
		assert.Equal(t, 42, value)
	})

	t.Run("returns env value when set and valid", func(t *testing.T) {
		key := "TEST_INT_" + t.Name()
		os.Setenv(key, "123")
		defer os.Unsetenv(key)

		value := getEnvAsInt(key, 42)
		assert.Equal(t, 123, value)
	})

	t.Run("returns default when env value is invalid", func(t *testing.T) {
		key := "TEST_INT_INVALID_" + t.Name()
		os.Setenv(key, "not_a_number")
		defer os.Unsetenv(key)

		value := getEnvAsInt(key, 42)
		assert.Equal(t, 42, value)
	})
}

func TestGetEnvAsDuration(t *testing.T) {
	t.Parallel()

	t.Run("returns default when env not set", func(t *testing.T) {
		value := getEnvAsDuration("NON_EXISTENT_DURATION", 5*time.Second)
		assert.Equal(t, 5*time.Second, value)
	})

	t.Run("returns env value when set and valid", func(t *testing.T) {
		key := "TEST_DURATION_" + t.Name()
		os.Setenv(key, "10s")
		defer os.Unsetenv(key)

		value := getEnvAsDuration(key, 5*time.Second)
		assert.Equal(t, 10*time.Second, value)
	})

	t.Run("returns default when env value is invalid", func(t *testing.T) {
		key := "TEST_DURATION_INVALID_" + t.Name()
		os.Setenv(key, "not_a_duration")
		defer os.Unsetenv(key)

		value := getEnvAsDuration(key, 5*time.Second)
		assert.Equal(t, 5*time.Second, value)
	})
}

// setupTestEnv sets up a clean environment for testing
func setupTestEnv(t *testing.T, env map[string]string) func() {
	t.Helper()

	// Save original environment
	original := make(map[string]string)
	for key := range env {
		original[key] = os.Getenv(key)
	}

	// Set test environment
	for key, value := range env {
		os.Setenv(key, value)
	}

	// Return cleanup function
	return func() {
		for key, value := range original {
			if value == "" {
				os.Unsetenv(key)
			} else {
				os.Setenv(key, value)
			}
		}
	}
}
