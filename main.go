package main

import (
	"fmt"
	"sort"
)

func median(list []int) float32 {
	// Sort the list
	sort.Ints(list)

	// If the length of the array is odd
	if len(list)%2 == 1 {
		return float32(list[len(list)/2])
	} else {
		return float32((list[len(list)/2] + list[len(list)/2-1])) / 2
	}
}

func fraudNotif(d int, expenditure []int) int {
	// Init counter
	count := 0

	// Iterate from d to len(expenditure)
	for i := d; i < len(expenditure); i++ {
		// Get the median
		median := median(expenditure[i-d : i])
		// If the current expenditure is greater than the median
		if float32(expenditure[i]) >= median*2 {
			// Increment the counter
			count++
		}
	}

	// Return Count
	return count
}

func main() {
	fmt.Println("Hello Perqara")
	fmt.Println("Fraud Notification Test")
	fmt.Println("Fraud Test Case 1: ", fraudNotif(3, []int{10, 20, 30, 40, 50}))
	fmt.Println("Fraud Test Case 2: ", fraudNotif(5, []int{2, 3, 4, 2, 3, 6, 8, 4, 5}))
	fmt.Println("Fraud Test Case 3: ", fraudNotif(4, []int{1, 2, 3, 4, 4}))
}
