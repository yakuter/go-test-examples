package main

import (
	"os"
	"testing"
)

func Sum(a, b int) int {
	return a + b
}

func setupTestCase(t *testing.T) func(t *testing.T) {
	t.Log("setup test case")
	return func(t *testing.T) {
		t.Log("teardown test case")
	}
}

func setupSubTest(t *testing.T) func(t *testing.T) {
	t.Log("setup sub test")
	return func(t *testing.T) {
		t.Log("teardown sub test")
	}
}

func TestAddition(t *testing.T) {
	cases := []struct {
		name     string
		a        int
		b        int
		expected int
	}{
		{"add", 2, 2, 4},
		{"minus", 0, -2, -2},
		{"zero", 0, 0, 0},
	}

	teardownTestCase := setupTestCase(t)
	defer teardownTestCase(t)

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			teardownSubTest := setupSubTest(t)
			defer teardownSubTest(t)

			result := Sum(tc.a, tc.b)
			if result != tc.expected {
				t.Fatalf("expected sum %v, but got %v", tc.expected, result)
			}
		})
	}
}

func TestDBFeatureA(t *testing.T) {
	defer models.TestDBManager.Reset()

	// Do the tests
}
func TestDBFeatureB(t *testing.T) {
	defer models.TestDBManager.Reset()

	// Do the tests
}
func TestMain(m *testing.M) {
	models.TestDBManager.Setup()
	// os.Exit() does not respect defer statements
	code := m.Run()
	models.TestDBManager.Exit()
	os.Exit(code)
}
