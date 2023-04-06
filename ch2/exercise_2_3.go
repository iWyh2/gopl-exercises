package ch2

var pc [256]byte

func init() {
	for i := range pc {
		pc[i] = pc[i/2] + byte(i&1)
	}
}

func PopCount(x uint64) int {
	var totalValue byte
	for i := 0; i < 8; i++ {
		totalValue += pc[byte(x>>(i*8))]
	}
	return int(totalValue)
}
