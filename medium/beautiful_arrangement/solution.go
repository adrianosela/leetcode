package beautiful_arrangement

import "math/bits"

func countArrangement(N int) int {
	possible := make([]uint, N+1)
	for i := 1; i <= N; i++ {
		for j := 1; j <= N; j++ {
			if i%j == 0 || j%i == 0 {
				possible[i] |= 1 << uint(j)
			}
		}
	}
	return backtrack(N, 0, possible)
}

func backtrack(pos int, used uint, possible []uint) int {
	if pos == 0 {
		return 1
	}

	total := 0
	options := possible[pos] & ^used
	for options != 0 {
		n := uint(bits.TrailingZeros(options))
		options &^= (1 << n)
		total += backtrack(pos-1, used|(1<<n), possible)
	}
	return total
}
