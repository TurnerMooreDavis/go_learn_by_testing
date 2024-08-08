package iteration

import (
	"fmt"
	"go-by-example/structs"
)

func Repeat(s string, r int) (repeatedString string) {
	repeatedString = s
	for i := 0; i < r; i++ {
		repeatedString += s
	}
	return
}
func ImpStruct() structs.Rectangle {
	foo := structs.Rectangle{Width: 10, Height: 10}
	fmt.Printf("foo: %v\n", foo)
	return foo
}
