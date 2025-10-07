package main

import (
	"errors"
)

func validateStatus(status string) error {
	if len(status) == 0 {
		return (errors.New("status cannot be empty"))
	}else if len(status) > 140 {
		return (errors.New("status exceeds 140 characters"))
	}

	return errors.New("")
}
