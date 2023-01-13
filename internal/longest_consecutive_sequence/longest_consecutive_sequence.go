package longest_consecutive_sequence

import "github.com/ashtanko/go-algorithms/utils"

// 128. Longest Consecutive Sequence
// https://leetcode.com/problems/longest-consecutive-sequence/
func longestConsecutive(nums []int) int {
	// Construct a set out of the nums array.
	numsSet := make(map[int]struct{})
	for _, n := range nums {
		numsSet[n] = struct{}{}
	}

	// The answer is stored here.
	maxSequenceLen := 0

	// Iterate through the set.
	for n := range numsSet {
		// We check if n-1 is in the set. If it is, then n is not the beginning of a sequence,
		// and we go to the next number immediately.
		if _, ok := numsSet[n-1]; !ok {
			// Otherwise, we increment n in a loop to see if the next consecutive value is stored in nums.
			seqLen := 1
			for {
				if _, ok = numsSet[n+seqLen]; ok {
					seqLen++
					continue
				}
				// When the sequence is over, see if we did better than before.
				maxSequenceLen = utils.Max(seqLen, maxSequenceLen)
				break
			}
		}
	}

	return maxSequenceLen
}
