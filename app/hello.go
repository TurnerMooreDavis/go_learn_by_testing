package main

import "fmt"

const suffix = " SAY WHAT!?!"
const reverseResponse = "YOU GOT ME "
const reverse = "no u"

func main() {
	fmt.Println(SayWhat("hello", ""))
}

func SayWhat(s string, s2 string) string {
	if s2 == reverse {
		return reverseResponse + s + "!"
	} else if s == "" {
		return "Dummy" + suffix
	}
	return s + suffix
}
