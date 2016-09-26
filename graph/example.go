package graph

// Returns a graph of 3 states and 3 transitions.
func TestGraph1() Explicit {
	// 0 -> [1, 2]
	// 1 -> [2]
	testFrom := []int{0, 2, 3, 3}
	testTo := []int{1, 2, 2}

	return Explicit{From: testFrom, To: testTo}
}

// Returns a graph of 11 states and 19 transitions.
func TestGraph2() Explicit {
	testFrom := []int{0, 2, 4, 5, 8, 8, 10, 12, 14, 17, 18, 19}
	testTo := []int{1, 7, 2, 5, 3, 1, 4, 5, 6, 9, 4, 10, 0, 8, 1, 7, 9, 10, 5}

	return Explicit{From: testFrom, To: testTo}
}
