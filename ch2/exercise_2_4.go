package ch2

func PopCountModified(x uint64) int {
	var res uint64
	for x != 0 {
		res += x & 1
		x >>= 1
	}
	return int(res)
}
