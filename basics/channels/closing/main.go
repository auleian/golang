package main

func countReports(numSentCh chan int) int {
	var counter int
	for {
		v, ok := <- numSentCh
		if !ok {
			break
		}
		counter += v
	}
	return counter
}

// don't touch below this line

func sendReports(numBatches int, ch chan int) {
	for i := range numBatches {
		numReports := i*23 + 32%17
		ch <- numReports
	}
	close(ch)
}
