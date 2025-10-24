package main

import "errors"

func deleteIfNecessary(users map[string]user, name string) (deleted bool, err error) {
	deleted = false
	person, ok := users[name]
	if !ok {
		err := errors.New("not found")
		return deleted, err
	}
	
	if person.scheduledForDeletion {
		delete(users, name)
		deleted = true
		return deleted, nil
	}else{
		return deleted, nil
	}
}

type user struct {
	name                 string
	number               int
	scheduledForDeletion bool
}
