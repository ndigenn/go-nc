package main

import (
	"flag"
	"fmt"
	"io"
	"log"
	"net"
	"os"
)

/*
TODO:
1. Define Command-Line Arguments

2. Create flags for server mode (-l), host (-h), and port (-p).
Parse Command-Line Arguments

3. Parse the flags and check if the -l flag is set to launch the server.
Implement the Server Logic

4. Set up a TCP listener using the specified host and port.
Accept incoming connections in a loop.
Handle each connection in a separate goroutine.
Process Client Connections

5. In the goroutine, read data from the connection and write it to stdout.
Close the connection when done.
Implement the Client Logic

6. Check if all required arguments are provided.
Connect to the server using the specified host and port.
Read data from stdin and send it to the server.
Handle Errors

7. Ensure proper error handling throughout the server and client logic.
Test the Server and Client

8. Run the server and client separately.
Use the standard netcat utility to verify the server's behavior.
Test the client by sending data to the server and ensuring it is received correctly.
*/


func processConnections(c net.Conn){
	_, err := io.Copy(os.Stdout, c)
	if (err != nil) {
		fmt.Printf("%s\n", err)
	}

	c.Close()
}

func startServer() {
	addr := fmt.Sprintf("%s:%d", *host, *port)
	listener, err := net.Listen("tcp", addr)
	if (err != nil){
		panic("error with server")
	}

	fmt.Printf("Listening for any connections: %s\n", listener.Addr().String())
	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Printf("Error accepting connections from the server: %s", err)
		} else {
			go processConnections(conn)
		}
	}
}

var (
	listen = flag.Bool("l", false, "listen")
	port = flag.Int("p", 0, "port")
	host = flag.String("h", "localhost", "host")
)

func main(){
	flag.Parse()
	if *listen {
		startServer()
		return
	}
}
