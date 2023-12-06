package total_money

import "github.com/ashtanko/go-algorithms/utils"

// https://leetcode.com/problems/calculate-money-in-leetcode-bank/
// totalMoneySimulate simulates the total amount of money accumulated over a given number of weeks.
// It calculates the total money by incrementing the amount each day starting from Monday.
// The function assumes that each Monday, the amount starts from 1 and increases by the day of the week.
// The simulation continues for the specified number of days (n).
// Parameters:
//   - n: The total number of days for the simulation.
//
// Returns:
//   - int: The total amount of money accumulated during the simulation.
func totalMoneySimulate(n int) int {
	// Initialize the total amount of money and the starting day of the week (Monday).
	ans := 0
	monday := 1

	// Continue the simulation while there are still days left.
	for n > 0 {
		// Iterate over the current week (up to 7 days or the remaining days if less than 7).
		for day := 0; day < utils.Min(n, 7); day++ {
			// Accumulate money for each day of the week.
			ans += monday + day
		}

		// Move to the next week and reduce the remaining days accordingly.
		n -= 7
		monday++
	}

	// Return the total amount of money accumulated during the simulation.
	return ans
}

const daysInWeek = 7
const initialMoney = 28

// totalMoneyMath calculates the total amount of money accumulated over a given number of weeks using a mathematical formula.
// It leverages an arithmetic progression formula to compute the sum of money for complete weeks,
// and then adds the sum for the remaining days in the final week.
// Parameters:
//   - n: The total number of days for which the money is calculated.
//
// Returns:
//   - int: The total amount of money accumulated based on the mathematical formula.
func totalMoneyMath(n int) int {
	// Calculate the number of complete weeks and the sum of money for the complete weeks.
	weeks := n / daysInWeek
	firstWeekSum := weeks * (2*initialMoney + (weeks-1)*daysInWeek) / 2

	// Calculate the number of remaining days after complete weeks.
	remainingDays := n % daysInWeek

	// Calculate the starting day of the final week.
	monday := 1 + weeks

	// Calculate the sum of money for the remaining days in the final week.
	finalWeekSum := 0
	for day := 0; day < remainingDays; day++ {
		finalWeekSum += monday + day
	}

	// Return the total amount of money accumulated based on the mathematical formula.
	return firstWeekSum + finalWeekSum
}
