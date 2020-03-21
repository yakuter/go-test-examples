package tests

import (
	"os"
	"testing"
)

func TestA(t *testing.T) {
	t.Log("Test A çalıştı")
}

func TestB(t *testing.T) {
	t.Log("Test B çalıştı")
}

func setup() {
	// Test için hazırlık
}
func teardown() {
	// Testten sonra temizlenecekler
}

func TestMain(m *testing.M) {
	setup()
	exitVal := m.Run()
	if exitVal == 0 {
		teardown()
	}
	os.Exit(exitVal)
}
