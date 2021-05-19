package main

import (
	"bufio"
	"fmt"
	"net"
	"strings"
)


type client struct {
	conn net.Conn
	name string
	room *room
	commands chan<- command
	admin bool
}

//function to read the input strings and figure out which command to use
func (c *client) readInput(){
	for{
		msg, err := bufio.NewReader(c.conn).ReadString('\n')
		//if theres an error exit the client read input
		if err != nil{
			return
		}

		//trims message of \r\n
		msg = strings.Trim(msg, "\r\n")
		//Splits message according to spaces
		args := strings.Split(msg, " ")
		//cmd = first word in input string
		cmd := strings.TrimSpace(args[0])

		switch cmd{
		case "help":
			c.commands <- command{
				id: CMD_HELP,
				client: c,
				args: args,
			}
		case "name":
			c.commands <- command{
				id: CMD_NAME,
				client: c,
				args: args,
			}
		case "join":
			c.commands <- command{
				id: CMD_JOIN,
				client: c,
				args: args,
			}
		case "rooms":
			c.commands <- command{
				id: CMD_ROOMS,
				client: c,
				args: args,
			}
		case "msg":
			c.commands <- command{
				id: CMD_MSG,
				client: c,
				args: args,
			}
		case "quit":
			c.commands <- command{
				id: CMD_QUIT,
				client: c,
				args: args,
			}
		case "motd":
			c.commands <- command{
				id: CMD_MOTD,
				client: c,
				args: args,
			}
		case "admin":
			c.commands <- command{
				id: CMD_ADM,
				client: c,
				args: args,
			}
		default:
			c.err(fmt.Errorf("unknown command: %s. Enter \"help\" for a list of commands", cmd))
		}
	}
}

func (c *client) err(err error){
	c.conn.Write([]byte("Error: " + err.Error() + "\n"))
}

func (c *client) msg(msg string){
	c.conn.Write([]byte("> " + msg + "\n"))
}
