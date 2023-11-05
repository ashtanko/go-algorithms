package advantage_shuffle

import (
	"sort"
)

// advantageCount rearranges the elements of slice A to maximize the number of elements in A that are greater
// than the corresponding elements in slice B. It returns a new slice with the rearranged elements of A.
func advantageCount(A []int, B []int) []int {
	// Create a sorted copy of slice A
	sortedA := make([]int, len(A))
	copy(sortedA, A)
	sort.Ints(sortedA)

	// Create a sorted copy of slice B
	sortedB := make([]int, len(B))
	copy(sortedB, B)
	sort.Ints(sortedB)

	// Create a map to store assigned elements for each element in B
	assigned := make(map[int][]int)

	// Create a slice to store remaining elements from A
	remaining := make([]int, 0)

	j := 0 // Index for sortedB

	// Iterate through the sortedA and assign elements to sortedB elements or store them as remaining
	for _, a := range sortedA {
		if a > sortedB[j] {
			assigned[sortedB[j]] = append(assigned[sortedB[j]], a)
			j++
		} else {
			remaining = append(remaining, a)
		}
	}

	// Create a result slice to store the final rearranged elements
	ans := make([]int, len(B))

	// Populate the result slice by assigning assigned values first and then remaining values
	for i, b := range B {
		if len(assigned[b]) > 0 {
			ans[i] = assigned[b][0]
			assigned[b] = assigned[b][1:]
		} else {
			ans[i] = remaining[0]
			remaining = remaining[1:]
		}
	}

	return ans
}
