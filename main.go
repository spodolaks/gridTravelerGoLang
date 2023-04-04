package main

import (
	"fmt"
	"time"
)

func gridTraveler(n int64, m int64) int64 {
	if n == 1 && m == 1{
		return 1
	}
	if (n == 0 || m == 0){
		return 0
	}

	return gridTraveler(n-1,m) + gridTraveler(n,m-1)
}

func main() {
	var n int64 = 16
	var m int64 = 16

	// memo := make(map[int64]int64)

	// Start the timer
	startTime := time.Now()

	result := gridTraveler(n, m)

	// Calculate the elapsed time in milliseconds
	elapsedTimeMs := time.Since(startTime).Milliseconds()
//testing
	fmt.Printf("For %d by %d grid number of ways: %d \n", n,m,result)
	fmt.Printf("Execution time: %d ms\n", elapsedTimeMs)
}
