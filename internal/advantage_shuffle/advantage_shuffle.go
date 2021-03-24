package advantage_shuffle

import (
	"github.com/ashtanko/go-algorithms/ds/deque"
	"sort"
)

func advantageCount(A []int, B []int) []int {

	sortedA := make([]int, len(A))
	copy(sortedA, A)
	sortedB := make([]int, len(B))
	copy(sortedB, B)
	sort.Ints(sortedA)
	sort.Ints(sortedB)

	assigned := make(map[int]*deque.Deque)
	for _, b := range B {
		assigned[b] = deque.NewDeque()
	}

	remaining := deque.NewDeque()
	j := 0
	for _, a := range sortedA {
		if a > sortedB[j] {
			assigned[sortedB[j]].Inject(a)
			j++
		} else {
			remaining.Inject(a)
		}
	}

	ans := make([]int, len(B))
	for i := 0; i < len(B); i++ {
		bDeq := assigned[B[i]]
		if bDeq.Size() > 0 {
			ans[i] = bDeq.Pop().(int)
		} else {
			ans[i] = remaining.Pop().(int)
		}
	}
	return ans
}
