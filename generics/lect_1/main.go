package main

func getLast[T any](s []T) T {
	// ?
	if len(s) == 0 {
		var zero T
		return zero
	}else{
		index := len(s)-1
		last := s[index]
		return last
	}
}
