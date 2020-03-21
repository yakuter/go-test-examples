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

func TestMain(m *testing.M) {
	// <setup code>
	exitVal := m.Run()
	if exitVal == 0 {
		// <tear-down code>
	}
	os.Exit(exitVal)
}
