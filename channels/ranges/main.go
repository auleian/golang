package main

func concurrentFib(n int) []int {
	// ?
	ch := make(chan int)
	values := []int{}

	go fibonacci(n, ch)

	for value:= range ch {
		values = append(values, value)
	}

	return values
}

// don't touch below this line

func fibonacci(n int, ch chan int) {
	x, y := 0, 1
	for i := 0; i < n; i++ {
		ch <- x
		x, y = y, x+y
	}
	close(ch)
}
