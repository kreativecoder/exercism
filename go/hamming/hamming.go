package hamming

import "errors"

//Distance : calculates the hamming distance between strings a and b
func Distance(a, b string) (int, error) {
	result := 0
	if len(a) == 0 || len(b) == 0 {
		return 0, nil
	}

	if len(a) == len(b) {
		for i := 0; i < len(a); i++ {
			if a[i] != b[i] {
				result++
			}
		}

		return result, nil
	}
	return result, errors.New("error raised")
}
