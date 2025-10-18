package main

func updateCounts(messagedUsers []string, validUsers map[string]int) {
	for i := range messagedUsers {
		if _, ok := validUsers[messagedUsers[i]]; ok {
			validUsers[messagedUsers[i]]++
		}
	}
	// increment count for each messaged user if they exist in validUsers
	// for _, user := range messagedUsers {
	// 	if _, ok := validUsers[user]; ok {
	// 		validUsers[user]++
	// 	}
	// }
}
