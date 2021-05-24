package module01

// Sum will sum up all of the numbers passed
// in and return the result
func Sum(numbers []int) int {
	sum := 0
	if numbers == nil || len(numbers) == 0 {
		return sum
	}
	for _, number := range numbers {
		sum += number
	}
	return sum
}
