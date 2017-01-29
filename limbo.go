package main

import (
	"log"
	"net"
	"bufio"
	"time"
	"math/rand"
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
	go KeepAlive()
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

func KeepAlive() {
	r := rand.New(rand.NewSource(15768735131534))
	keepalive := &PacketPlayKeepAlive{
		id: 0,
	}
	for {
		for _, player := range players {
			if player.state == PLAY {
				if player.keepalive != 0 {
					player.Kick("Timed out")
				}

				id := int(r.Uint32())
				if id == 0 {
					id = 1
				}
				keepalive.id = id
				player.keepalive = id
				player.WritePacket(keepalive)
			}
		}
		time.Sleep(3000000000)
	}
}

func HandleConnection(conn net.Conn, id int) {
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
		keepalive: 0,
	}
	player.register(id)

	for {
		packet, err := player.ReadPacket()
		if err != nil {
			break
		}

		CallEvent("packetReceived", packet)
	}

	player.unregister(id)
	conn.Close()
	log.Printf("%s disconnected.", conn.RemoteAddr().String())
}