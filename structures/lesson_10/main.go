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
	var person User
	person.Name = name
	person.Type = membershipType

	if membershipType == "premium" {
		person.MessageCharLimit = 1000
	}else {
		person.MessageCharLimit = 100
	}

	return person
}
