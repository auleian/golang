package main

func countConnections(groupSize int) int {
	totalConnections := 0
	for i := 1; i <= groupSize - 1; i++ {
		totalConnections += i
	}

	return totalConnections
}
