package module01

// BaseToDec takes in a number and the base it is currently
// in and returns the decimal equivalent as an integer.
//
// Eg:
//
//   BaseToDec("E", 16) => 14
//   BaseToDec("1110", 2) => 14
//
func basePow(base int, pow int) int {
	product := 1

	for i := 0; i < pow; i++ {
		product = product * base
	}

	return product
}

func Char2Decimal(c rune) int {
	A := 'A'
	Zero := '0'

	if c >= '0' && c <= '9' {
		return int(c) - int(Zero)
	} else {
		return (int(c) - int(A)) + 10 // A is 10
	}
}
func BaseToDec(value string, base int) int {
	sum := 0
	maxPow := len(value) - 1
	for i, digit := range value {
		// var value int
		// fmt.Sscanf(string(digit), "%X", &value)
		// sum += value * basePow(base, (maxPow-i))
		sum += Char2Decimal(digit) * basePow(base, (maxPow-i))
	}

	return sum
}
