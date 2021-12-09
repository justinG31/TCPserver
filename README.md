# TCPserver
Justin Gomez, Sebastian Fernandez, Gabriella Munger

Description
----
This project simulates a Network Denial of Service through a method of a TCP SYN flood attack. The prrogram simulates the attack by the flood.go acting as the adversary performing the TCP SYN flood attack. The file flood.go asks a user to input the amount of clients to send to the server as well as the amount of GOMAXPROCS() to utilize. The server, server.go, asks the user for the amount of GOMAXPROCS() to utilize than starts running the server. Once the server is running, it will run until receiving and EOF interrupt cammand i.e. "CRTL-C" or when an abundance of clients can succesfully trigger the server to stop running. 


Input and Output
----
## server.go Input
The input will ask the user to specify the amount of GOMAXPROCS() to use

## server.go Output
The output will reflect the amount of clients that have connected to the server and the time each client connected since the start of the server running.

## flood.go Input
The input will consist of the amount of clients to send to the target server as well as the amount of GOMAXPROCS() to utilize when running.

## flood.go Output
currently there is no output

How to Run:
----
clone the repository : ' git clone https://github.com/justinG31/TCPserver.git '


## Run server
go run server.go

## Run client(s)
go run flood.go



