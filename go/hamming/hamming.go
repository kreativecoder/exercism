package hamming

import "errors"

//Distance : calculates the hamming distance between strings a and b
func Distance(a, b string) (int, error) {
	result := 0

	if len(a) == len(b) {
		for i := range a {
			if a[i] != b[i] {
				result++
			}
		}

		return result, nil
	}

	return -1, errors.New("error raised")
}
