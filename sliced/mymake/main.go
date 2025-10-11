package main

func getMessageCosts(messages []string) []float64 {
	// ?
	cost := make([]float64, len(messages))

	for i:= range messages{
		cost[i] = float64(len(messages[i])) * 0.01
	}

	return cost
}
