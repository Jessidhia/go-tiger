package tiger

import (
	"unsafe"
)

// Uses unsafe to pun the c uint64 into an array of its bytes
// Overall benchmark 110.49 MB/s
func (t *tiger) tigerRound(a, b, c int, x, mul uint64) {
	t.state[c] ^= x
	cb := [8]byte(*(*[8]byte)(unsafe.Pointer(&t.state[c])))

	t.state[a] -= sBox[0][cb[0]] ^ sBox[1][cb[2]] ^ sBox[2][cb[4]] ^ sBox[3][cb[6]]
	t.state[b] += sBox[3][cb[1]] ^ sBox[2][cb[3]] ^ sBox[1][cb[5]] ^ sBox[0][cb[7]]
	t.state[b] *= mul
}

// Does the round in the proper, mathy way
// Overall benchmark 105.44 MB/s
func (t *tiger) tigerRound_math(a, b, c int, x, mul uint64) {
	t.state[c] ^= x
	tc := t.state[c]
	cb := [8]byte{
		byte((tc >> (0 * 8)) & 0xff), byte((tc >> (1 * 8)) & 0xff),
		byte((tc >> (2 * 8)) & 0xff), byte((tc >> (3 * 8)) & 0xff),
		byte((tc >> (4 * 8)) & 0xff), byte((tc >> (5 * 8)) & 0xff),
		byte((tc >> (6 * 8)) & 0xff), byte((tc >> (7 * 8)) & 0xff),
	}

	t.state[a] -= sBox[0][cb[0]] ^ sBox[1][cb[2]] ^ sBox[2][cb[4]] ^ sBox[3][cb[6]]
	t.state[b] += sBox[3][cb[1]] ^ sBox[2][cb[3]] ^ sBox[1][cb[5]] ^ sBox[0][cb[7]]
	t.state[b] *= mul
}
