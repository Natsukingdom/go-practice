package fizzbuzz

import (
	"strconv"
)

func FizzBuzzResult(n int) []string {
	results := make([]string, n)
	for i := 1; i <= n; i++ {
		results = append(results, judge_fizzbuzz(i))
	}
	return results
}

func judge_fizzbuzz(n int) string {
	switch {
	case n%15 == 0:
		return "FizzBuzz"
	case n%5 == 0:
		return "Buzz"
	case n%3 == 0:
		return "Fizz"
	default:
		return strconv.Itoa(n)
	}
}
