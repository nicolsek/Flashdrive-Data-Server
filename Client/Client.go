package main

import (
	"fmt"
	"net"
	"os"
	"runtime"
	"strconv"
	"time"
)

// main ... The main function?
func main() {
	client := new(Client)
	server := createServer("localhost", 8080, "tcp")

	dialServer(server, getClientData(client))
}

//Data regarding the client, will send to the server
type Client struct {
	time     string
	OS       string
	cpuCount int
	hostName string
}

//Data regarding how to connect to the server
type Server struct {
	IPv4           string
	ConnectionType string
	PortD          int
}

// getClientData ... Gets the data to transmit to the server and sets it for the client
func getClientData(client *Client) []byte {
	//Seperator to reference the end of a particular part of data
	sep := "\n"
	//Finding the time and date
	client.time = time.Now().String()
	//Finding the OS
	client.OS = runtime.GOOS
	//Finding the hostName
	hostName, _ := os.Hostname()
	client.hostName = hostName
	//Finding the cpu count
	client.cpuCount = runtime.NumCPU()

	data := make([]byte, 512)

	//Creates a single string for all the data
	dataString := client.time + sep + client.OS + sep + client.hostName + sep + strconv.Itoa(client.cpuCount)

	fmt.Printf("Data: \n%v", dataString)

	//Take the slice of that into the data array
	copy(data[:], dataString)

	return data
}

// createServer ... Creates a server with the properties of a server and returns it
func createServer(IPv4 string, PortD int, ConnectionType string) *Server {
	server := new(Server)
	//Sets the IPv4
	server.IPv4 = IPv4
	//Sets the Port
	server.PortD = PortD
	//Sets the connectionType
	server.ConnectionType = ConnectionType

	return server
}

// dialServer ... Dials the server and connects to it, will send information and then close the connection
func dialServer(server *Server, data []byte) {
	conn, _ := net.Dial(server.ConnectionType, string(server.IPv4)+":"+string(server.PortD))
	conn.Write(data)
	defer conn.Close()
}
