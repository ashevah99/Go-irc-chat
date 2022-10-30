# Go IRC chat
Just a little project to learn Golang by making an IRC chat server. 

Goals/overview: 
IRC server: 
- relay messages asynchronously 
- multiple users 


IRC chat: 
- send & receive chats to server 
- broadcast messages 

Stretch: 
- different protocols to send/receive data 
- cache chat logs 


# Usage
Server:  go run main.go simpleClient.go simpleServer.go -mode server -port 6969
Client: go run * -mode client -port 6969 -ip 127.0.0.1
# Progress  
10/30/22: 
- got a simple client / server model working 
- command line argument parsing (flag library) 

# TODOS
- setup launch.json to compile multiple golang files 
- get multiple clients to server 
- eventually tackle bigger, more dynamic client & server 
- fiddle with cooler ways of sending data (i.e. covert channels, other protocols (dns, udp, whatever irc is,  etc.)) 
- add a GUI 