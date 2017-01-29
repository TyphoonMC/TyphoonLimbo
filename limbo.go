package main

import (
	"log"
	"net"
	"bufio"
)

var (
	connCounter = 0
)

func main() {
	InitPackets()

	ln, err := net.Listen("tcp", ":25565") //TODO config file for port definition
	if err != nil {
		log.Fatal(err)
	}
	log.Println("Server launched.")
	for {
		conn, err := ln.Accept()
		if err != nil {
			log.Print(err)
		} else {
			connCounter+=1
			go HandleConnection(conn, connCounter)
		}
	}
}

func HandleConnection(conn net.Conn, uid int) {
	log.Printf("%s connected.", conn.RemoteAddr().String())

	player := &Player {
		conn: conn,
		state: HANDSHAKING,
		protocol: V1_10,
		io: &ConnReadWrite{
			rdr: bufio.NewReader(conn),
			wtr: bufio.NewWriter(conn),
		},
		inaddr: InAddr{
			"",
			0,
		},
		name: "",
		uuid: "d979912c-bb24-4f23-a6ac-c32985a1e5d3",
	}

	for {
		packet, err := player.ReadPacket()
		if err != nil {
			break
		}

		CallEvent("packetReceived", packet)
	}

	conn.Close()
	log.Printf("%s disconnected.", conn.RemoteAddr().String())
}