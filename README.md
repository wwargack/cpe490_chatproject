# cpe490_chatproject
CPE 490 Final Project. Build a client/server chat application using Go.

-------------------------------------------------------------------------------

Start the server by running "go run ." inside the directory of all of the files.

Connect to the server using "telnet (address) 8888" where (address) is the ip address of the server.

For example (address) is: localhost on the same machine, or 192.168.1.27 on my home network since my PC is at that address.

-------------------------------------------------------------------------------

Commands:

help 	       - lists available commands

name (string)  - changes display name to (string)

join (string)  - joins room names (string), creates room named (string) if (string) doesn't exist yet

rooms	       - lists available rooms

msg (string)   - broadcasts (string) as a message to everyone connected

quit           - disconnect from the server

admin (string) - changes the user to an admin if (string) is the correct password 

motd (string)  - changes the room's motd to (string), only usable by admins