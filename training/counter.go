package main

func makeCounter(start, step int) func() int {
	
	return func () int{
		current := start
		start += step
		return current
	}
}

func composeInt(f func(int) int, g func(int) int) func(int) int {
	return func (x int) int{
		return f(g(x))
	}
}
