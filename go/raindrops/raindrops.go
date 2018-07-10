package raindrops

import "fmt"

//Convert : accepts parameter num and return string representation based on factors
func Convert(num int) string {
	result := ""

	if num%3 == 0 {
		result += "Pling"
	}

	if num%5 == 0 {
		result += "Plang"
	}

	if num%7 == 0 {
		result += "Plong"
	}

	if len(result) == 0 {
		result = fmt.Sprintf("%d", num)
	}
	return result
}
