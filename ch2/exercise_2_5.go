package ch2

func PopCountModified2(x uint64) int {
	var res uint64
	for x != 0 {
		x = x & (x - 1)
		res++
	}
	return int(res)
}
