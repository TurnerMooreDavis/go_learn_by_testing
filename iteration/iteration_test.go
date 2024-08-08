package iteration

import (
	"go-by-example/structs"
	"go-by-example/utils"
	"testing"
)

func TestRepeat(t *testing.T) {
	got := Repeat("a", 4)
	want := "aaaaa"
	if got != want {
		t.Errorf("got %q, want %q", got, want)
	}
}
func BenchmarkRepeat(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Repeat("a", 4)
	}
}

func TestImp(t *testing.T) {
	rect := ImpStruct()
	got := structs.RectArea(rect)
	var want float32 = 100
	utils.Assert(t, got, want)
}
