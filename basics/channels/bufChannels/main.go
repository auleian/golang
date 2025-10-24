package main

func addEmailsToQueue(emails []string) chan string {
	// ?
	ch := make(chan string, len(emails))

	for i := range emails{
		ch <- emails[i]
	}

	return ch
}
