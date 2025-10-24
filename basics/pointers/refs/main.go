package main

import (
	"strings"
)

func removeProfanity(message *string) {
	badwords := []string{"fubb", "shiz", "witch"}
	for _, bad := range badwords {
		asterisk := strings.Repeat("*", len(bad))
		*message = strings.ReplaceAll(*message, bad, asterisk)
	}
	
}
