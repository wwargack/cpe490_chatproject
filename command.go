package main

type commandID int

//declaration of command ids: help, name, message, users, quit
const (
	CMD_HELP commandID = iota
	CMD_NAME
	CMD_JOIN
	CMD_ROOMS
	CMD_MSG
	CMD_QUIT
	CMD_MOTD
	CMD_ADM
)

//creates command structure with values id, pointer to the client, string args
type command struct {
	id     commandID
	client *client
	args   []string
}