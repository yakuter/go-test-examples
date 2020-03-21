package tests

import (
	"testing"
)

func TestParalel(t *testing.T) {
	t.Parallel() // marks TLog as capable of running in parallel with other tests
	tests := []struct {
		name string
	}{
		{"test 1"},
		{"test 2"},
		{"test 3"},
		{"test 4"},
	}
	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			t.Error(tt.name)
		})
	}
}
