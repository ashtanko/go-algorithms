package power_of_four

import "math"

func isPowerOfFourBruteForce(n int) bool {
	if n == 0 {
		return false
	}
	if n == 1 {
		return true
	}
	if n%4 != 0 {
		return false
	}

	return isPowerOfFourBruteForce(n / 4)
}

func isPowerOfFourLogarithmic(n int) bool {
	if n == 0 {
		return false
	}

	result := math.Log(float64(n)) / math.Log(4)
	return result == math.Floor(result)
}

// isPowerOfFourBitwise checks if a given integer is a power of four using bitwise operations.
// It returns true if the input is a power of four, otherwise, it returns false.
func isPowerOfFourBitwise(n int) bool {
	if n <= 0 {
		return false
	}
	if n == 1 {
		return true
	}

	// Check if the least significant bit (LSB) position is even and if n is a power of two
	if Lsb(n)%2 == 0 && (n&(n-1)) == 0 {
		return true
	} else {
		return false
	}
}

// Lsb calculates the position of the least significant bit (LSB) set in the binary representation of an integer.
// It returns the position (0-based) of the LSB if one is set; otherwise, it returns -1.
func Lsb(n int) int {
	p := 1
	for n > 0 {
		if n&1 > 0 {
			return p - 1
		}
		p++
		n = n >> 1
	}
	return -1
}
