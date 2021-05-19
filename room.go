package main

import "net"

type room struct{
	name string
	members map[net.Addr]*client
	motd string
}

//function to broadcast message to everyone in the room besides sender
func (r *room) broadcast(sender *client, msg string){
	for addr, m := range r.members{
		if addr != sender.conn.RemoteAddr(){
			m.msg(msg)
		}
	}
}