package module01

// Reverse will return the provided word in reverse
// order. Eg:
//
//   Reverse("cat") => "tac"
//   Reverse("alphabet") => "tebahpla"
//
func Reverse(word string) string {
	var rWord string
	for _, c := range word {
		rWord = string(c) + rWord
	}
	return rWord
}
