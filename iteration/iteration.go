package iteration

func Repeat(s string, r int) (repeatedString string) {
	repeatedString = s
	for i := 0; i < r; i++ {
		repeatedString += s
	}
	return
}
