package main

import "testing"

func TestHello(t *testing.T) {
	t.Run("name say what", func(t *testing.T) {
		var got string = SayWhat("Turner", "")
		want := "Turner" + suffix
		assert(t, got, want)
	})
	t.Run("if no name given say 'Dummy SAY WHAT!?!'", func(t *testing.T) {
		var got string = SayWhat("", "")
		want := "Dummy" + suffix
		assert(t, got, want)
	})

	t.Run("if given a second argument 'no u' and first argument 'Turner' say 'YOU GOT ME Turner!'", func(t *testing.T) {
		var got string = SayWhat("Turner", "no u")
		want := "YOU GOT ME Turner!"
		assert(t, got, want)
	})
}

func assert(t testing.TB, got, want string) {
	t.Helper()
	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}
