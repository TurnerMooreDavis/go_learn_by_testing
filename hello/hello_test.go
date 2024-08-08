package main

import (
	"go-by-example/utils"
	"testing"
)

func TestHello(t *testing.T) {
	t.Run("name say what", func(t *testing.T) {
		var got string = SayWhat("Turner", "")
		want := "Turner" + defaultSuffix
		utils.Assert(t, got, want)
	})
	t.Run("if no name given say 'Dummy SAY WHAT!?!'", func(t *testing.T) {
		var got string = SayWhat("", "")
		want := "Dummy" + defaultSuffix
		utils.Assert(t, got, want)
	})

	t.Run("if given a second argument 'no u' and first argument 'Turner' say  'Turner GOT ME?!? AHHH!'", func(t *testing.T) {
		var got string = SayWhat("Turner", "no u")
		want := "Turner" + reverseSuffix
		utils.Assert(t, got, want)
	})
	t.Run("if given a second argument 'wha...' and first argument 'Turner' say 'Turner YOUR SO DUMB HAHAHAhHHAH'", func(t *testing.T) {
		var got string = SayWhat("Turner", "wha...")
		want := "Turner" + whaSuffix
		utils.Assert(t, got, want)
	})
}
