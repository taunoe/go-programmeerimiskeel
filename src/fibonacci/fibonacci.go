package main

func fib_one(position int) int {
	seq := []int{0, 1} // starts with 0, 1, 1, 2, 3, 5, 8

	for len(seq) <= position {
		// Arvuta uus jada kuni jõuame soovitud positsioonini
		seq = append(seq, seq[len(seq)-1]+seq[len(seq)-2]) // Viimane + eelviimane
	}

	return seq[len(seq)-1] // viimane
}

// rekursioon
func fib_two(position int, seq []int) int {
	// First call
	if seq == nil {
		seq = []int{0, 1}
	}

	for len(seq) <= position {
		seq = append(seq, seq[len(seq)-1]+seq[len(seq)-2])
		return fib_two(position, seq)
	}

	return seq[len(seq)-1]
}

func fib_three(n int) int {
	if n < 2 {
		return n
	}

	a, b := fib_three(n-1), fib_three(n-2)
	return (a + b)
}
