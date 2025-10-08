package main

func maxMessages(thresh int) int {
totalCost := 0
	for i := 0; ; i++ {
		messageCost := 100 + i
		if totalCost + messageCost > thresh {
			return i
		}
		totalCost += messageCost
	}
}
