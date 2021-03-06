package main

import (
	"fmt"
	"io/ioutil"
	"net"
	"os"
	"strings"
)

// main ... Main function?
func main() {
	server := createServer("Server", "tcp", "localhost:8080")
	debugServer(server)
	startServer(server)
}

// Server ... Is the type Server with the properties of Name, Ipv4, ConnectionType, PortL, and isReading. Simple object reference to a server.
type Server struct {
	Name string
	//IPv4           string
	ConnectionType string
	PortL          string
	isReading      bool
	isActive       bool
}

// createServer ... Creates a server and returns *Server is non-active.
func createServer(Name string, ConnectionType string, PortL string) *Server {
	server := new(Server)

	//Set the name to the name
	server.Name = Name
	//Get the IPv4 by using the built-in function
	//server.IPv4 = ipnet.IP.string()
	//Configure the connection type by either TCP, UDP, or FTP
	server.ConnectionType = ConnectionType
	//Specify the port that it should listen on
	server.PortL = PortL
	//Tell if the server is reading data
	server.isReading = false
	//Tell if the server is active and recieving data
	server.isActive = false

	return server
}

// startServer ... Starts the server and runs a listening for-loop, all actions should be done before this is called.
func startServer(server *Server) {
	//Using the listen function activate the server so that it can listen for incoming data
	ln, err := net.Listen(server.ConnectionType, server.PortL)
	//Check if the error is nil, if it is then write the error and exit
	if err != nil {
		os.Stderr.WriteString("Oops: " + err.Error() + "\n")
		os.Exit(1)
	}

	for {
		//Check connection and accept it
		conn, err := ln.Accept()

		if conn == nil {
			server.isReading = false
		}

		if err != nil {
			os.Stderr.WriteString("Oops: " + err.Error() + "\n")
			os.Exit(1)
		}

		if conn != nil {
			server.isReading = true
			//Read the data concurrently
			go readData(conn)
		}
	}

}

// debugServer ... Gives information about the server to see what it's doing and possibly to debug it using that information.
func debugServer(server *Server) {
	//Makes my life easier by just a tiny bit.
	nl := "\n"
	//Writing the console as this should be used as a console application.
	//This is to give information regarding the properties afformentioned about the server.
	fmt.Printf("Server Name: %v", server.Name+nl)
	// fmt.Printf("Server IPv4: %v", server.IPv4+nl)
	fmt.Printf("Server ConnectionType: %v", server.ConnectionType+nl)
	fmt.Printf("Server PortL: %v", server.PortL+nl)
	// fmt.Printf("Server isReading: %v", server.isReading+nl)
	// os.Stdout.Writestring("Server isActive: " + server.isActive + nl)
}

// readData ... Reads data from the incoming go connection and stores it into a byte array and closes the connection.
func readData(conn net.Conn) {
	data := make([]byte, 512)
	data, _ = ioutil.ReadAll(conn)

	writeData(conn, data)
}

// writeData ... Writes data from the conn and data to a file with the name of the remote address and the data from the connection.
func writeData(conn net.Conn, data []byte) {
	//Seperates the port from the RemoteAddr and just takes the IP
	ioutil.WriteFile(strings.Split(conn.RemoteAddr().String(), ":")[0], data, 0644)
}
