package main

import "fmt"

const defaultSuffix = " SAY WHAT!?!"
const reverseSuffix = " GOT ME?!? AHHH!"
const whaSuffix = " YOUR SO DUMB HAHAHAhHHAH"

const reverse = "no u"
const wha = "wha..."
const defaultName = "Dummy"

func main() {
	fmt.Println(SayWhat("hello", ""))
}

func getSuffix(response string) (suffix string) {
	switch response {
	case reverse:
		suffix = reverseSuffix
	case wha:
		suffix = whaSuffix
	default:
		suffix = defaultSuffix
	}
	return
}

func getName(nameInput string) (name string) {
	if nameInput == "" {
		name = defaultName
	} else {
		name = nameInput
	}
	return
}

func SayWhat(nameInput string, response string) string {
	return getName(nameInput) + getSuffix(response)
}
