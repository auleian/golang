package main

import (
	"errors"
)

const (
	planFree = "free"
	planPro  = "pro"
)

func getMessageWithRetriesForPlan(plan string, messages [3]string) ([]string, error) {
	var err error = errors.New("unsupported plan")
	
	switch plan {
		case planPro:
			mymessages := messages[:]
			return mymessages, nil
		case planFree:
			mymessages := messages[:2]
			return mymessages, nil
		default:
			mymessages := messages[:0]
			return mymessages, err
	}
}
