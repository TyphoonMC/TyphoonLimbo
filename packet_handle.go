package main

import "log"

type PacketHandshake struct {
	protocol Protocol
	address string
	port uint16
	state State
}
func (packet *PacketHandshake) Read(player *Player) (err error) {
	protocol, err := player.ReadVarInt()
	if err != nil {
		log.Print(err)
		return
	}
	packet.protocol = Protocol(protocol)
	packet.address, err = player.ReadString()
	if err != nil {
		log.Print(err)
		return
	}
	packet.port, err = player.ReadUInt16()
	if err != nil {
		log.Print(err)
		return
	}
	state, err := player.ReadVarInt()
	if err != nil {
		log.Print(err)
		return
	}
	packet.state = State(state)
	return
}
func (packet *PacketHandshake) Write(player *Player) (err error) {
	return
}
func (packet *PacketHandshake) Handle(player *Player) {
	player.state = packet.state
	player.protocol = packet.protocol
	player.inaddr.address = packet.address
	player.inaddr.port = packet.port
	return
}
func (packet *PacketHandshake) Id() int {
	return 0x00
}

type PacketStatusRequest struct {}
func (packet *PacketStatusRequest) Read(player *Player) (err error) {
	return
}
func (packet *PacketStatusRequest) Write(player *Player) (err error) {
	return
}
func (packet *PacketStatusRequest) Handle(player *Player) {
	response := PacketStatusResponse{
		response: `{"version":{"name":"1.10","protocol":210},"players":{"max":10000,"online":0,"sample":[]},"description":{"text":"TyphoonLimbo v0.1"},"favicon":"data:image/png;base64,<data>"}`,
	}
	player.WritePacket(response)
	return
}
func (packet *PacketStatusRequest) Id() int {
	return 0x00
}

type PacketStatusResponse struct {
	response string
}
func (packet *PacketStatusResponse) Read(player *Player) (err error) {
	return
}
func (packet *PacketStatusResponse) Write(player *Player) (err error) {
	err = player.WriteString(packet.response)
	if err != nil {
		log.Print(err)
		return
	}
	return
}
func (packet *PacketStatusResponse) Handle(player *Player) {
	return
}
func (packet *PacketStatusResponse) Id() int {
	return 0x00
}
