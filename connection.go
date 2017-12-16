package main

import (
	"net"
	"log"
	"compress/zlib"
	"bytes"
	"encoding/binary"
)

type State int8
const (
	HANDSHAKING State = iota
	STATUS
	LOGIN
	PLAY
)

type Gamemode uint8
const (
	SURVIVAL Gamemode = iota
	CREATIVE
	ADVENTURE
	SPECTATOR
)

type Dimension uint32
const (
	NETHER Dimension = 0xFF
	OVERWORLD Dimension = 0
	END Dimension = 1
)

type Difficulty uint8
const (
	PEACEFUL Difficulty = iota
	EASY
	NORMAL
	HARD
)

type ChatPosition uint8
const (
	CHAT_BOX ChatPosition = iota
	SYSTEM
	ACTION_BAR
)

type ScoreboardPosition uint8
const (
	LIST ScoreboardPosition = iota
	SIDEBAR
	BELOW_NAME
)

type BossBarAction int
const (
	BOSSBAR_ADD BossBarAction = iota
	BOSSBAR_REMOVE
	BOSSBAR_UPDATE_HEALTH
	BOSSBAR_UPDATE_TITLE
	BOSSBAR_UPDATE_STYLE
	BOSSBAR_UPDATE_FLAGS
)

type BossBarColor int
const (
	BOSSBAR_COLOR_PINK BossBarColor = iota
	BOSSBAR_COLOR_BLUE
	BOSSBAR_COLOR_RED
	BOSSBAR_COLOR_GREEN
	BOSSBAR_COLOR_YELLOW
	BOSSBAR_COLOR_PURPLE
	BOSSBAR_COLOR_WHITE
)

type BossBarDivision int
const (
	BOSSBAR_NODIVISION BossBarDivision = iota
	BOSSBAR_6NOTCHES
	BOSSBAR_10NOTCHES
	BOSSBAR_12NOTCHES
	BOSSBAR_20NOTCHES
)

type LevelType string
const (
	DEFAULT LevelType = "default"
	FLAT LevelType = "flat"
	LARGE_BIOMES LevelType = "largeBiomes"
	AMPLIFIED LevelType = "amplified"
	DEFAULT_1_1 LevelType = "default_1_1"
)

type ChunkSection struct {
	bits_per_block uint8
	palette_length int
	palette []int
	data_array_length int
	data_array []uint64
	block_light []uint8
	sky_light []uint8
}

type Protocol uint16
const (
	V1_7_2 Protocol = 4
	V1_7_6 = 5
	V1_8 = 47
	V1_9 = 107
	V1_9_1 = 108
	V1_9_2 = 109
	V1_9_3 = 110
	V1_10 = 210
	V1_11 = 315
	V1_11_1 = 316
	V1_12 = 335
	V1_12_1 = 338
	V1_12_2 = 340
)
var (
	COMPATIBLE_PROTO = [...]Protocol{
		V1_7_2, V1_7_6,
		V1_8,
		V1_9, V1_9_1, V1_9_2, V1_9_3,
		V1_10,
		V1_11, V1_11_1,
		V1_12, V1_12_1,
		V1_12_2,
	}
)

func IsCompatible(proto Protocol) bool {
	for _, x := range COMPATIBLE_PROTO {
		if x == proto {
			return true
		}
	}
	return false
}

type InAddr struct {
	address string
	port uint16
}

type Player struct {
	id int
	conn net.Conn
	io *ConnReadWrite
	state State
	protocol Protocol
	inaddr InAddr
	name string
	uuid string
	keepalive int
	compression bool
}

func (player *Player) ReadPacket() (packet Packet, err error){
	if !player.compression {
		return player.ReadPacketWithoutCompression()
	} else {
		return player.ReadPacketWithCompression()
	}
}

func (player *Player) ReadPacketWithoutCompression() (packet Packet, err error) {
	length, err := player.ReadVarInt()
	if err != nil {
		return
	}

	id, err := player.ReadVarInt()
	if err != nil {
		return
	}
	if player.state == PLAY {
		id = player.HackServerbound(id)
	}

	packet, err = player.HandlePacket(id, length)
	if err != nil {
		return
	} else if packet != nil {
		if config["logs"].(bool) {
			log.Println("->", id, packet)
		}
		packet.Handle(player)
	}
	return
}

func (player *Player) ReadPacketWithCompression() (packet Packet, err error) {
	packetLength, err := player.ReadVarInt()
	if err != nil {
		return
	}

	dataLength, err := player.ReadVarInt()
	if err != nil {
		return
	}
	dataLengthLength := binary.PutUvarint(player.io.buffer[:], uint64(dataLength))

	var id int
	var length int

	if dataLength == 0 {
		id, err = player.ReadVarInt()
		if err != nil {
			return
		}
		idLength := binary.PutUvarint(player.io.buffer[:], uint64(id))
		length = packetLength-dataLengthLength-idLength

		if player.state == PLAY {
			id = player.HackServerbound(id)
		}
		packet, err = player.HandlePacket(id, length)

		if err != nil {
			return
		} else if packet != nil {
			if config["logs"].(bool) {
				log.Println("u->", id, packet)
			}
			packet.Handle(player)
		}
	} else {
		var compressed []byte = make([]byte, packetLength-dataLengthLength)
			player.conn.Read(compressed[:])

		r, err := zlib.NewReader(bytes.NewReader(compressed[:]))
		if err != nil {
			return nil, err
		}

		var uncompressed []byte = make([]byte, packetLength-dataLengthLength)
		r.Read(uncompressed[:])
		r.Close()

		tmp := player.io
		rdr := bytes.NewBuffer(uncompressed[:])
		player.io.rdr = rdr

		id, err = player.ReadVarInt()
		if err != nil {
			return nil, err
		}

		if player.state == PLAY {
			id = player.HackServerbound(id)
		}
		packet, err = player.HandlePacket(id, length)
		player.io = tmp

		if err != nil {
			return nil, err
		} else if packet != nil {
			if config["logs"].(bool) {
				log.Println("c->", id, packet)
			}
			packet.Handle(player)
		}
	}
	return
}

func (player *Player) WritePacket(packet Packet) (err error){
	if !player.compression {
		return player.WritePacketWithoutCompression(packet)
	} else {
		return player.WritePacketWithCompression(packet)
	}
}

func (player *Player) WritePacketWithoutCompression(packet Packet) (err error){
	buff := newVarBuffer(256)
	tmp := player.io
	player.io = &ConnReadWrite{
		rdr: tmp.rdr,
		wtr: buff,
	}

	id := packet.Id()
	if player.state == PLAY {
		id = player.HackClientbound(id)
	}
	if id == -1 {
		return
	}
	player.WriteVarInt(id)
	packet.Write(player)

	ln := newVarBuffer(0)
	player.io.wtr = ln
	player.WriteVarInt(buff.Len())
	player.io = tmp
	player.conn.Write(ln.Bytes())
	player.conn.Write(buff.Bytes())

	if config["logs"].(bool) {
		log.Println("<-", id, packet)
	}
	return nil
}

func (player *Player) WritePacketWithCompression(packet Packet) (err error){
	buff := newVarBuffer(256)
	tmp := player.io
	player.io = &ConnReadWrite{
		rdr: tmp.rdr,
		wtr: buff,
	}

	id := packet.Id()
	if player.state == PLAY {
		id = player.HackClientbound(id)
	}
	if id == -1 {
		return
	}
	player.WriteVarInt(id)
	packet.Write(player)

	var rBuff []byte;
	var dataLength = 0
	if buff.Len() < compressionThreshold {
		rBuff = buff.Bytes()
	} else {
		var b bytes.Buffer
		w := zlib.NewWriter(&b)
		w.Write(buff.Bytes())
		w.Close()
		rBuff = b.Bytes()
		dataLength = len(rBuff)
	}

	buff2 := newVarBuffer(1)
	player.io.wtr = buff2
	player.WriteVarInt(dataLength)
	packetLength := len(rBuff)+buff2.Len()

	buff3 := newVarBuffer(1)
	player.io.wtr = buff3
	player.WriteVarInt(packetLength)

	player.io = tmp
	player.conn.Write(buff3.Bytes())
	player.conn.Write(buff2.Bytes())
	player.conn.Write(rBuff)

	if config["logs"].(bool) {
		if buff.Len() < compressionThreshold {
			log.Println("<-u", id, packet)
		} else {
			log.Println("<-c", id, packet)
		}
	}
	return nil
}
