package tests

import "testing"

func Hello(name string) string {
	return "Hello, " + name
}

func TestHello(t *testing.T) {

	assertCorrectMessage := func(t *testing.T, got, want string) {
		t.Helper()
		if got != want {
			t.Errorf("got %q want %q", got, want)
		}
	}

	t.Run("saying hello to Erhan", func(t *testing.T) {
		got := Hello("Erhan")
		want := "Hello, Erhan"
		assertCorrectMessage(t, got, want)
	})

	t.Run("saying hello to Mehmet", func(t *testing.T) {
		got := Hello("Mehmet")
		want := "Hello, Mehmet"
		assertCorrectMessage(t, got, want)
	})

}
