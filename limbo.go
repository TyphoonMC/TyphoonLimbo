package main

import (
	"log"
	"net"
	"bufio"
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
			go HandleConnection(conn)
		}
	}
}

func HandleConnection(conn net.Conn) {
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