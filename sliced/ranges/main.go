package main

import "slices"

func indexOfFirstBadWord(msg []string, badWords []string) int {
	for  i := range msg {
		if slices.Contains(badWords, msg[i]) {
				return i
			}
	}

	return -1
}
