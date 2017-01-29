package main

import (
	"log"
	"net"
	"bufio"
)

func main() {
	initPackets()

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
			go handleConnection(conn)
		}
	}
}

func handleConnection(conn net.Conn) {
	log.Printf("%s connected.", conn.RemoteAddr().String())

	player := &Player {
		conn: conn,
		state: HANDSHAKING,
		protocol: V1_10,
		io: ConnReadWrite{
			rdr: bufio.NewReader(conn),
			wtr: bufio.NewWriter(conn),
		},
		inaddr: InAddr{
			"",
			0,
		},
	}

	packet, err := player.readPacket()
	if err != nil {
		conn.Close()
		return
	}

	callEvent("packetReceived", packet)
}