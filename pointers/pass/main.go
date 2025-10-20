package main

type Analytics struct {
	MessagesTotal     int
	MessagesFailed    int
	MessagesSucceeded int
}

type Message struct {
	Recipient string
	Success   bool
}

// don't touch above this line

// ?
func analyzeMessage (a *Analytics, m Message){
	if m.Success {
		a.MessagesSucceeded += 1
	}else{
		a.MessagesFailed += 1
	}
	a.MessagesTotal += 1
}