package tests

import (
	"testing"
)

// Go ile Subtest mantığıyla test yazma
func TestSubtests(t *testing.T) {
	// <setup code>
	t.Run("A", func(t *testing.T) {
		t.Log("Test A1 tamamlandı")
	})
	t.Run("B", func(t *testing.T) {
		t.Log("Test B tamamlandı")
	})
	t.Run("C", func(t *testing.T) {
		t.Log("Test C tamamlandı")
	})
	// <tear-down code>
}
