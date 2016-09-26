package graph

func TestGraph1() Explicit {
	// 0 -> 1
	// 0 -> 2
	// 1 -> 2
	testFrom := []int{0, 2, 3, 3}
	testTo   := []int{1, 2, 2}

	return Explicit{From: testFrom, To: testTo}
}