package main


func bulkSend(numMessages int) float64 {
	totalCost := 0.0
	for i := 0.0; i < float64(numMessages); i ++{
		totalCost += 1.0 + i*0.01
	}

	return totalCost
}
