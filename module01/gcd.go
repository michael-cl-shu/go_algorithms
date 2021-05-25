package module01

// Given two numbers, A and B:

// Step 1: If B == 0, return A
// Step 2: A becomes B, and B becomes the remainder of dividing A by B
// `a, b = b, a % b`
// Step 3: Go to step 1

func GCD(a, b int) int {
	// if b == 0 {
	// 	return a
	// } else {
	// 	a, b = b, a%b
	// 	return GCD(a, b)
	// }
	for b != 0 {
		a, b = b, a%b
	}
	return a
}
