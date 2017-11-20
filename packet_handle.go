package main

import (
	"log"
	"fmt"
	"strings"
	"github.com/satori/go.uuid"
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
	protocol := COMPATIBLE_PROTO[0]
	if IsCompatible(player.protocol) {
		protocol = player.protocol
	}

	max_players := int(config["max_players"].(float64))
	motd := config["motd"].(string)

	if max_players < players_count && !config["restricted"].(bool) {
		max_players = players_count
	}

	response := PacketStatusResponse{
		response: fmt.Sprintf(`{"version":{"name":"Typhoon","protocol":%d},"players":{"max":%d,"online":%d,"sample":[]},"description":{"text":"%s"},"favicon":""}`, protocol, max_players, players_count, motd),
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
var(
	join_game = PacketPlayJoinGame{
		entity_id: 0,
		gamemode: SPECTATOR,
		dimension: END,
		difficulty: NORMAL,
		level_type: DEFAULT,
		max_players: 0xFF,
		reduced_debug: false,

	}
	position_look = PacketPlayerPositionLook{}
)
func (packet *PacketLoginStart) Handle(player *Player) {
	if !IsCompatible(player.protocol) {
		player.LoginKick("Incompatible version")
		return
	}

	max_players := int(config["max_players"].(float64))

	if max_players <= players_count && config["restricted"].(bool) {
		player.LoginKick("Server is full")
	}

	player.name = packet.username

	success := PacketLoginSuccess{
		uuid: player.uuid,
		username: player.name,
	}
	player.WritePacket(&success)
	player.state = PLAY
	player.register()

	player.WritePacket(&join_game)
	player.WritePacket(&position_look)

	if &join_message != nil {
		player.WritePacket(&join_message)
	}
	if &bossbar_create != nil {
		player.WritePacket(&bossbar_create)
	}
	//player.Kick("Not implemented yet..")
	return
}
func (packet *PacketLoginStart) Id() int {
	return 0x00
}

type PacketLoginDisconnect struct {
	component string
}
func (packet *PacketLoginDisconnect) Read(player *Player) (err error) {
	return
}
func (packet *PacketLoginDisconnect) Write(player *Player) (err error) {
	err = player.WriteString(packet.component)
	if err != nil {
		log.Print(err)
		return
	}
	return
}
func (packet *PacketLoginDisconnect) Handle(player *Player) {
	return
}
func (packet *PacketLoginDisconnect) Id() int {
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

type PacketPlayChat struct {
	message string
}
func (packet *PacketPlayChat) Read(player *Player) (err error) {
	packet.message, err = player.ReadString()
	if err != nil {
		log.Print(err)
		return
	}
	return
}
func (packet *PacketPlayChat) Write(player *Player) (err error) {
	return
}
func (packet *PacketPlayChat) Handle(player *Player) {
	if len(packet.message) > 0 && packet.message[0] != '/' {
		player.WritePacket(&PacketPlayMessage{
			component: fmt.Sprintf(`{"text":"%s: %s"}`, player.name, strings.Replace(packet.message, `"`, `\"`, -1)),
			position: CHAT_BOX,
		})
	}
	return
}
func (packet *PacketPlayChat) Id() int {
	return 0x02
}

type PacketPlayMessage struct {
	component string
	position ChatPosition
}
func (packet *PacketPlayMessage) Read(player *Player) (err error) {
	return
}
func (packet *PacketPlayMessage) Write(player *Player) (err error) {
	err = player.WriteString(packet.component)
	if err != nil {
		log.Print(err)
		return
	}
	err = player.WriteUInt8(uint8(packet.position))
	if err != nil {
		log.Print(err)
		return
	}
	return
}
func (packet *PacketPlayMessage) Handle(player *Player) {
	return
}
func (packet *PacketPlayMessage) Id() int {
	return 0x0F
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

type PacketPlayChunkData struct {
	x uint32
	y uint32
	ground_up_continuous bool
	primary_bit_mask int
	data_size int
	data []ChunkSection
	biomes [256]byte
	block_entities_length int
}
func (packet *PacketPlayChunkData) Read(player *Player) (err error) {
	return
}
func (packet *PacketPlayChunkData) Write(player *Player) (err error) {
	return
}
func (packet *PacketPlayChunkData) Handle(player *Player) {
	return
}
func (packet *PacketPlayChunkData) Id() int {
	return 0x20
}

type PacketPlayJoinGame struct {
	entity_id uint32
	gamemode Gamemode
	dimension Dimension
	difficulty Difficulty
	max_players uint8
	level_type LevelType
	reduced_debug bool
}
func (packet *PacketPlayJoinGame) Read(player *Player) (err error) {
	return
}
func (packet *PacketPlayJoinGame) Write(player *Player) (err error) {
	err = player.WriteUInt32(packet.entity_id)
	if err != nil {
		log.Print(err)
		return
	}
	err = player.WriteUInt8(uint8(packet.gamemode))
	if err != nil {
		log.Print(err)
		return
	}
	err = player.WriteUInt32(uint32(packet.dimension))
	if err != nil {
		log.Print(err)
		return
	}
	err = player.WriteUInt8(uint8(packet.difficulty))
	if err != nil {
		log.Print(err)
		return
	}
	err = player.WriteUInt8(packet.max_players)
	if err != nil {
		log.Print(err)
		return
	}
	err = player.WriteString(string(packet.level_type))
	if err != nil {
		log.Print(err)
		return
	}
	err = player.WriteBool(packet.reduced_debug)
	if err != nil {
		log.Print(err)
		return
	}
	return
}
func (packet *PacketPlayJoinGame) Handle(player *Player) {
	return
}
func (packet *PacketPlayJoinGame) Id() int {
	return 0x23
}

type PacketPlayerPositionLook struct {
	x float64
	y float64
	z float64
	yaw float32
	pitch float32
	flags uint8
	teleport_id int
}
func (packet *PacketPlayerPositionLook) Read(player *Player) (err error) {
	return
}
func (packet *PacketPlayerPositionLook) Write(player *Player) (err error) {
	err = player.WriteFloat64(packet.x)
	if err != nil {
		log.Print(err)
		return
	}
	err = player.WriteFloat64(packet.y)
	if err != nil {
		log.Print(err)
		return
	}
	err = player.WriteFloat64(packet.z)
	if err != nil {
		log.Print(err)
		return
	}
	err = player.WriteFloat32(packet.yaw)
	if err != nil {
		log.Print(err)
		return
	}
	err = player.WriteFloat32(packet.pitch)
	if err != nil {
		log.Print(err)
		return
	}
	err = player.WriteUInt8(packet.flags)
	if err != nil {
		log.Print(err)
		return
	}
	err = player.WriteVarInt(packet.teleport_id)
	if err != nil {
		log.Print(err)
		return
	}
	return
}
func (packet *PacketPlayerPositionLook) Handle(player *Player) {
	return
}
func (packet *PacketPlayerPositionLook) Id() int {
	return 0x2E
}

type PacketBossBar struct {
	uuid uuid.UUID
	action BossBarAction
	title string
	health float32
	color BossBarColor
	division BossBarDivision
	flags uint8
}
func (packet *PacketBossBar) Read(player *Player) (err error) {
	return
}
func (packet *PacketBossBar) Write(player *Player) (err error) {
	err = player.WriteUUID(packet.uuid)
	if err != nil {
		log.Print(err)
		return
	}
	err = player.WriteVarInt(int(packet.action))
	if err != nil {
		log.Print(err)
		return
	}
	if packet.action == BOSSBAR_UPDATE_TITLE || packet.action == BOSSBAR_ADD {
		err = player.WriteString(packet.title)
		if err != nil {
			log.Print(err)
			return
		}
	}
	if packet.action == BOSSBAR_UPDATE_HEALTH || packet.action == BOSSBAR_ADD {
		err = player.WriteFloat32(packet.health)
		if err != nil {
			log.Print(err)
			return
		}
	}
	if packet.action == BOSSBAR_UPDATE_STYLE || packet.action == BOSSBAR_ADD {
		err = player.WriteVarInt(int(packet.color))
		if err != nil {
			log.Print(err)
			return
		}
		err = player.WriteVarInt(int(packet.division))
		if err != nil {
			log.Print(err)
			return
		}
	}
	if packet.action == BOSSBAR_UPDATE_STYLE || packet.action == BOSSBAR_ADD {
		err = player.WriteUInt8(packet.flags)
		if err != nil {
			log.Print(err)
			return
		}
	}
	return
}
func (packet *PacketBossBar) Handle(player *Player) {
	return
}
func (packet *PacketBossBar) Id() int {
	return 0x0C
}
