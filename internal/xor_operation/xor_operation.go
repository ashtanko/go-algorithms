package xor_operation

func xorOperation(n int, start int) int {
	res := start
	for i := 1; i < n; i++ {
		res = res ^ (start + 2*i)
	}
	return res
}
