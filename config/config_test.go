package config

import (
	"github.com/google/go-cmp/cmp"
	"os"
	"testing"
)

type testConfig struct {
	Name     string
	Budget   float64
	IsTest   bool
	Database struct {
		Name string
		Port int
	}
	Items  []string
	Secret EnvVariable
}

func TestConfig_LoadFromWithProfile(t *testing.T) {
	tests := []struct {
		name         string
		profile      string
		envVariables map[string]string
		expected     testConfig
	}{
		{
			name: "Default case",
			expected: testConfig{
				Name:   "Awesome Service",
				Budget: 47.11,
				IsTest: true,
				Database: struct {
					Name string
					Port int
				}{
					Name: "testDB",
					Port: 4711,
				},
				Items:  []string{"Item 1", "Item 2", "Item 3"},
				Secret: "P@SSW0RD",
			},
			envVariables: map[string]string{
				"SECRET": "P@SSW0RD",
			},
		},
		{
			name:    "With profile",
			profile: "staging",
			expected: testConfig{
				Name:   "Another Service",
				Budget: -15.11,
				IsTest: false,
				Database: struct {
					Name string
					Port int
				}{
					Name: "stagingDB",
					Port: 1511,
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			for key, value := range tt.envVariables {
				if err := os.Setenv(key, value); err != nil {
					t.Fatal(err)
				}
			}

			cfg := testConfig{}

			err := LoadFromWithProfile("test", tt.profile, &cfg)
			if err != nil {
				t.Fatalf("loadFromFile failed: %s", err.Error())
			}

			if diff := cmp.Diff(tt.expected, cfg); diff != "" {
				t.Fatalf(diff)
			}
		})
	}
}

func TestConfig_CurrentProfile(t *testing.T) {
	tests := []struct {
		name     string
		expected string
		envKey   string
	}{
		{
			name:     "No profile",
			expected: "",
		},
		{
			name:     "With profile",
			expected: "MY_PROFILE",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := os.Setenv(envKeyProfile, tt.expected); err != nil {
				t.Fatal(err)
			}

			actual := CurrentProfile()

			if diff := cmp.Diff(tt.expected, actual); diff != "" {
				t.Fatalf(diff)
			}
		})
	}
}
func TestConfig_MustGetEnv(t *testing.T) {
	tests := []struct {
		name      string
		setKey    bool
		wantError bool
	}{
		{
			name:      "Exist",
			setKey:    true,
			wantError: false,
		},
		{
			name:      "Panic",
			setKey:    false,
			wantError: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			defer func() {
				errorOccurred := recover() != nil
				if errorOccurred != tt.wantError {
					t.Fatalf("Expected error %v, got error %v", tt.wantError, errorOccurred)
				}
			}()
			if err := os.Unsetenv("TEST_KEY"); err != nil {
				t.Fatal(err)
			}
			if tt.setKey {
				if err := os.Setenv("TEST_KEY", "I am here..."); err != nil {
					t.Fatal(err)
				}
			}

			actual := MustGetEnv("TEST_KEY")

			if diff := cmp.Diff("I am here...", actual); diff != "" {
				t.Fatalf(diff)
			}
		})
	}
}
