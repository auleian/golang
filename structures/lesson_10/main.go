package main

type User struct {
	membership
	Name string
}

type membership struct {
	Type string
	MessageCharLimit int
}

func newUser(name string, membershipType string) User {
	// ?
	
}
