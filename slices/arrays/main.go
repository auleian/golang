package main

func getMessageWithRetries(primary, secondary, tertiary string) ([3]string, [3]int) {
	messages := [3]string{primary, secondary, tertiary}
	cost := [3]int{len(primary), (len(secondary)+len(primary)), (len(tertiary)+ len(secondary)+len(primary))}

	return messages, cost
}
