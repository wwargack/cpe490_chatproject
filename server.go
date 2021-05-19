package main

import (
	"fmt"
	"log"
	"net"
	"strings"
)

//declares server structure that houses a list of names and commands
type server struct{
	rooms map[string]*room
	commands chan command
}

//function new server creates a server
func newServer() *server {
	return &server{
		rooms: make(map[string]*room),
		commands: make(chan command),
	}
}

//runs the server. runs commands based on inputs
func (s *server) run(){
	for cmd:=range s.commands{
		switch cmd.id{
		case CMD_HELP:
			s.help(cmd.client, cmd.args)
		case CMD_NAME:
			s.setName(cmd.client, cmd.args)
		case CMD_JOIN:
			s.joinRoom(cmd.client, cmd.args)
		case CMD_ROOMS:
			s.getRooms(cmd.client, cmd.args)
		case CMD_MSG:
			s.sendMsg(cmd.client, cmd.args)
		case CMD_QUIT:
			s.quit(cmd.client, cmd.args)
		case CMD_ADM:
			s.loginAdmin(cmd.client, cmd.args)
		case CMD_MOTD:
			s.changeMOTD(cmd.client, cmd.args)
		}
	}
}

//function to create a new client with connection, name, and commands input
func (s *server) newClient(conn net.Conn){
	//logs new client connection with ip in the server
	log.Printf("A new client connected: %s", conn.RemoteAddr().String)

	c := &client{
		conn: conn,
		name: "username",
		commands: s.commands,
	}
	
	c.readInput()
}

//help command outputs a list of commands to the input client
func (s *server) help(c *client, args []string){
	c.msg("\"help\" - lists available commands \n")
	c.msg("\"name (string)\" - changes display name to (string) \n")
	c.msg("\"join (string)\" - joins room names (string), creates room named (string) if (string) doesn't exist yet \n")
	c.msg("\"rooms\" - lists available rooms \n")
	c.msg("\"msg (string)\" - broadcasts (string) as a message to everyone connected \n")
	c.msg("\"quit\" - disconnect from the server \n")
	c.msg("\"admin (string)\" - changes the user to an admin if (string) is the correct password \n")
	c.msg("\"motd (string)\" - changes the room's motd to (string), only usable by admins \n")
}

//changes user's name to input
func (s *server) setName(c *client, args []string){
	c.name = args[1]
	c.msg(fmt.Sprintf("Display name changed to: %s", c.name))
}

//broadcasts message into the room
func (s *server) sendMsg(c *client, args []string){
	c.room.broadcast(c, c.name + ": " + strings.Join(args[1:len(args)], " "))
}

func (s *server) joinRoom(c *client, args []string) {
	roomName := args[1]

	r, ok := s.rooms[roomName]
	if !ok {
		r = &room{
			name:    roomName,
			members: make(map[net.Addr]*client),
		}
		s.rooms[roomName] = r
	}
	r.members[c.conn.RemoteAddr()] = c
	s.leaveRoom(c)
	c.room = r

	r.broadcast(c, fmt.Sprintf("%s has joined the room.", c.name))
	c.msg(fmt.Sprintf("Welcome to %s", r.name))
	c.msg(fmt.Sprintf(r.motd))
}

func (s *server) getRooms(c *client, args []string) {
	var rooms []string
	for name := range s.rooms {
		rooms = append(rooms, name)
	}
	c.msg(fmt.Sprintf("Available rooms are: %s", strings.Join(rooms, ", ")))
}

//user quits the server
func (s *server) quit(c *client, args []string) {
	log.Printf("A client has disconnected: %s", c.conn.RemoteAddr().String())
	s.leaveRoom(c)
	c.conn.Close()
}

//changes the user to be an admin if they know the password
func (s *server) loginAdmin(c *client, args []string){
	if args[1] == "test123"{
		c.admin = true
		c.msg("You are now an admin.")
	} else{
		c.msg("Incorrect password.")
	}
}

func (s *server) changeMOTD(c *client, args []string){
	if c.admin{
		c.room.motd = strings.Join(args[1:len(args)], " ")
	}
}

//removes user from names list
func (s *server) leaveRoom(c *client){
	//if user joins a room
	if c.room != nil{
		delete(c.room.members, c.conn.RemoteAddr())
		c.room.broadcast(c, fmt.Sprintf("%s left the room.", c.name))
	}
}

