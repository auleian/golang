package main

import "errors"

func getUserMap(names []string, phoneNumbers []int) (map[string]user, error) {
	contacts := make(map[string]user)
	
	if len(names) < len(phoneNumbers) || len(names) > len(phoneNumbers) {
		err := errors.New("invalid sizes")
		return nil, err
	}

	for i := range names {
		contacts[names[i]] = user{name: names[i], phoneNumber: phoneNumbers[i],}
	}

	return contacts, nil
}

type user struct {
	name        string
	phoneNumber int
}
