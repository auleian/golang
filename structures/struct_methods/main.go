package main

type authenticationInfo struct {
	username string
	password string
}

func (auth authenticationInfo) getBasicAuth() string{
	output := "Authorization: Basic " +  auth.username + ":" + auth.password
	return output
}