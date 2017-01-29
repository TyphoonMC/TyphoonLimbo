package main

import (
	"log"
	"fmt"
)

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
	protocol := player.protocol

	response := PacketStatusResponse{
		response: fmt.Sprintf(`{"version":{"name":"Typhoon 1.10","protocol":%d},"players":{"max":10000,"online":%d,"sample":[]},"description":{"text":"TyphoonLimbo v0.1"},"favicon":""}`, protocol, len(players)),
	}
	player.WritePacket(&response)
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

type PacketStatusPing struct {
	time uint64
}
func (packet *PacketStatusPing) Read(player *Player) (err error) {
	packet.time, err = player.ReadUInt64()
	if err != nil {
		log.Print(err)
		return
	}
	return
}
func (packet *PacketStatusPing) Write(player *Player) (err error) {
	err = player.WriteUInt64(packet.time)
	if err != nil {
		log.Print(err)
		return
	}
	return
}
func (packet *PacketStatusPing) Handle(player *Player) {
	player.WritePacket(packet)
	return
}
func (packet *PacketStatusPing) Id() int {
	return 0x01
}

type PacketLoginStart struct {
	username string
}
func (packet *PacketLoginStart) Read(player *Player) (err error) {
	packet.username, err = player.ReadString()
	if err != nil {
		log.Print(err)
		return
	}
	return
}
func (packet *PacketLoginStart) Write(player *Player) (err error) {
	return
}
func (packet *PacketLoginStart) Handle(player *Player) {
	player.name = packet.username

	success := PacketLoginSuccess{
		uuid: player.uuid,
		username: player.name,
	}
	player.WritePacket(&success)
	player.state = PLAY
	player.register()

	//player.Kick("Not implemented yet..")
	return
}
func (packet *PacketLoginStart) Id() int {
	return 0x00
}

type PacketLoginSuccess struct {
	uuid string
	username string
}
func (packet *PacketLoginSuccess) Read(player *Player) (err error) {
	return
}
func (packet *PacketLoginSuccess) Write(player *Player) (err error) {
	err = player.WriteString(packet.uuid)
	if err != nil {
		log.Print(err)
		return
	}
	err = player.WriteString(packet.username)
	if err != nil {
		log.Print(err)
		return
	}
	return
}
func (packet *PacketLoginSuccess) Handle(player *Player) {
	return
}
func (packet *PacketLoginSuccess) Id() int {
	return 0x02
}

type PacketPlayDisconnect struct {
	component string
}
func (packet *PacketPlayDisconnect) Read(player *Player) (err error) {
	return
}
func (packet *PacketPlayDisconnect) Write(player *Player) (err error) {
	err = player.WriteString(packet.component)
	if err != nil {
		log.Print(err)
		return
	}
	return
}
func (packet *PacketPlayDisconnect) Handle(player *Player) {
	return
}
func (packet *PacketPlayDisconnect) Id() int {
	return 0x1A
}

type PacketPlayKeepAlive struct {
	id int
}
func (packet *PacketPlayKeepAlive) Read(player *Player) (err error) {
	packet.id, err = player.ReadVarInt()
	if err != nil {
		log.Print(err)
		return
	}
	return
}
func (packet *PacketPlayKeepAlive) Write(player *Player) (err error) {
	err = player.WriteVarInt(packet.id)
	if err != nil {
		log.Print(err)
		return
	}
	return
}
func (packet *PacketPlayKeepAlive) Handle(player *Player) {
	if player.keepalive != packet.id {
		player.Kick("Invalid keepalive")
	}
	player.keepalive = 0
	return
}
func (packet *PacketPlayKeepAlive) Id() int {
	return 0x1F
}