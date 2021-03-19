package keys_and_rooms

func canVisitAllRooms(rooms [][]int) bool {
	seen := make([]bool, len(rooms))
	seen[0] = true
	var stack []int
	stack = append(stack, 0)

	for len(stack) != 0 {
		index := len(stack) - 1
		element := (stack)[index]
		stack = (stack)[:index]
		for _, v := range rooms[element] {
			if !seen[v] {
				seen[v] = true
				stack = append(stack, v)
			}
		}
	}

	for _, v := range seen {
		if !v {
			return false
		}
	}

	return true
}
