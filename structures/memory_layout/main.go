package main

type contact struct {
	userID       string
	sendingLimit int32
	age          int32
}

type perms struct {
	canSend         bool
	canReceive      bool
	canManage       bool
	permissionLevel int
}
