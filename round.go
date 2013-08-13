package tiger

func (t *tiger) tigerPass(a, b, c int, x [8]uint64, mul uint64) {
	for i := 0; i < 8; i++ {
		t.tigerRound((a+i)%3, (b+i)%3, (c+i)%3, x[i], mul)
	}
}

func (t *tiger) tigerBlock(x [8]uint64) {
	// abc_store
	st := t.state

	// first pass
	t.tigerPass(0, 1, 2, x, 5)

	// second pass
	x = keySchedule(x)
	t.tigerPass(2, 0, 1, x, 7)

	// third pass
	x = keySchedule(x)
	t.tigerPass(1, 2, 0, x, 9)

	// feed_forward
	t.state[0] ^= st[0]
	t.state[1] -= st[1]
	t.state[2] += st[2]
}
