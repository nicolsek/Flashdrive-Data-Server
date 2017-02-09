package main 

import ( 

	"os"
	"net"
	"time"
	"runtime"

)

var {
	serverAddress := "192.168.1.1"
}

// main ... The main function?
func main()
	client Client = new(Client)
	server := createServer(serverAddress, )

	dialServer(server, getClientData())
}

//Data regarding the client, will send to the server
type Client struct {
	time time.Time
	OS string
	cpuCount int
	hostName string
}

//Data regarding how to connect to the server
type Server struct {
	IPv4 string
	ConnectionType string
	PortD int
}

// getClientData ... Gets the data to transmit to the server and sets it for the client
func getClientData(client *Client) *[]byte {
	//Seperator to reference the end of a particular part of data
	sep := "\n"
	//Finding the time and date
	client.time = Time.Now()
	//Finding the OS
	client.OS = runtime.GOOS
	//Finding the hostName
	hostName, _ := os.Hostname()
	client.hostName = hostName
	//Finding the cpu count
	client.cpuCount = runtime.NumCpu()

	data := make([]byte, 512)

	//Creates a single string for all the data
	dataString := client.time + sep + client.OS + sep + client.hostName + sep + client.cpuCount

	//Take the slice of that into the data array
	copy(data[:], dataString)

	return data
}

// createServer ... Creates a server with the properties of a server and returns it
func createServer(IPV4 string, PortD int, ConnectionType string) *Server {
	server Server = new(Server)
	//Sets the IPv4
	server.IPV4 = IPV4
	//Sets the Port
	server.PortD = PortD
	//Sets the connectionType
	server.ConnectionType = ConnectionType
}

// dialServer ... Dials the server and connects to it, will send information and then close the connection
func dialServer(server *Server, data *[]byte) {
	conn, _ := net.Dial(server.ConnectionType, "%v:%v", server.IPv4, server.PortD)
	conn.Write(data)
	defer conn.Close()
}

