package utils

import (
	"reflect"
	"testing"
)

func Assert(t testing.TB, got, want any) {
	t.Helper()
	if !reflect.DeepEqual(got, want) {
		t.Errorf("got '%v' want '%v'", got, want)
		wantType := reflect.TypeOf(want)
		gotType := reflect.TypeOf(got)
		if gotType != wantType {
			t.Errorf("types incompatible got: '%v' want: '%v'", gotType, wantType)
		}
	}
}
