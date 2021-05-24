package module01

// DecToBase will return a string representing
// the provided decimal number in the provided base.
// This is limited to bases 2-16 for simplicity.
//
// Eg:
//
//   DecToBase(14, 16) => "E"
//   DecToBase(14, 2) => "1110"
//var digits = []string{"0", "1", "2", "3"}

var digits = []string{"0", "1", "2", "3", "4", "5", "6", "7", "8", "9", "A", "B", "C", "D", "E", "F"}

func DecToBase(dec, base int) string {

	if dec > 0 {
		remainder := dec % base
		quotient := dec / base
		return DecToBase(quotient, base) + digits[remainder]
	} else {
		// dec == 0
		return ""
	}
}
