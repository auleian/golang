package main

import "fmt"

func (a *analytics) handleEmailBounce(em email) error {
	em.recipient.updateStatus(em.status)
	a.track(em.status)
	fmt.Println("Hello")
	return nil
}
