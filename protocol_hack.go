package main

var (
	clientbound = make(map[Protocol]map[int]int)
	serverbound = make(map[Protocol]map[int]int)
	HACKS       = [...]Protocol{V1_8, V1_12, V1_12_1, V1_12_2}
)

func InitHacks() {
	// Hack 1.8
	clientbound[V1_8] = make(map[int]int)
	clientbound[V1_8][0x00] = 0x0E
	clientbound[V1_8][0x01] = 0x11
	clientbound[V1_8][0x02] = 0x2C
	clientbound[V1_8][0x03] = 0x0F
	clientbound[V1_8][0x04] = 0x10
	clientbound[V1_8][0x05] = 0x0C
	clientbound[V1_8][0x06] = 0x0B
	clientbound[V1_8][0x07] = 0x37
	clientbound[V1_8][0x08] = 0x25
	clientbound[V1_8][0x09] = 0x35
	clientbound[V1_8][0x0A] = 0x24
	clientbound[V1_8][0x0B] = 0x23
	clientbound[V1_8][0x0C] = -1
	clientbound[V1_8][0x0D] = 0x41
	clientbound[V1_8][0x0E] = 0x3A
	clientbound[V1_8][0x0F] = 0x02
	clientbound[V1_8][0x10] = 0x22
	clientbound[V1_8][0x11] = 0x32
	clientbound[V1_8][0x12] = 0x2E
	clientbound[V1_8][0x13] = 0x2D
	clientbound[V1_8][0x14] = 0x30
	clientbound[V1_8][0x15] = 0x31
	clientbound[V1_8][0x16] = 0x2F
	clientbound[V1_8][0x17] = -1
	clientbound[V1_8][0x18] = 0x3F
	clientbound[V1_8][0x19] = 0x2F
	clientbound[V1_8][0x1A] = 0x40
	clientbound[V1_8][0x1B] = 0x1A
	clientbound[V1_8][0x1C] = 0x27
	clientbound[V1_8][0x1D] = -1
	clientbound[V1_8][0x1E] = 0x2B
	clientbound[V1_8][0x1F] = 0x00
	clientbound[V1_8][0x20] = 0x21
	clientbound[V1_8][0x21] = 0x28
	clientbound[V1_8][0x22] = 0x2A
	clientbound[V1_8][0x23] = 0x01
	clientbound[V1_8][0x24] = 0x34
	clientbound[V1_8][0x25] = 0x15
	clientbound[V1_8][0x26] = 0x17
	clientbound[V1_8][0x27] = 0x16
	clientbound[V1_8][0x28] = 0x14
	clientbound[V1_8][0x29] = -1
	clientbound[V1_8][0x2A] = 0x36
	clientbound[V1_8][0x2B] = 0x39
	clientbound[V1_8][0x2C] = 0x42
	clientbound[V1_8][0x2D] = 0x38
	clientbound[V1_8][0x2E] = 0x08
	clientbound[V1_8][0x2F] = 0x04
	clientbound[V1_8][0x30] = 0x13
	clientbound[V1_8][0x31] = 0x1E
	clientbound[V1_8][0x32] = 0x48
	clientbound[V1_8][0x33] = 0x07
	clientbound[V1_8][0x34] = 0x19
	clientbound[V1_8][0x35] = 0x44
	clientbound[V1_8][0x36] = 0x43
	clientbound[V1_8][0x37] = 0x09
	clientbound[V1_8][0x38] = 0x3D
	clientbound[V1_8][0x39] = 0x1C
	clientbound[V1_8][0x3A] = 0x1B
	clientbound[V1_8][0x3B] = 0x12
	clientbound[V1_8][0x3C] = 0x04
	clientbound[V1_8][0x3D] = 0x1F
	clientbound[V1_8][0x3E] = 0x06
	clientbound[V1_8][0x3F] = 0x3B
	clientbound[V1_8][0x40] = -1
	clientbound[V1_8][0x41] = 0x3E
	clientbound[V1_8][0x42] = 0x3C
	clientbound[V1_8][0x43] = 0x05
	clientbound[V1_8][0x44] = 0x03
	clientbound[V1_8][0x45] = 0x45
	clientbound[V1_8][0x46] = 0x33
	clientbound[V1_8][0x47] = -1
	clientbound[V1_8][0x48] = 0x47
	clientbound[V1_8][0x49] = 0x0D
	clientbound[V1_8][0x4A] = 0x18
	clientbound[V1_8][0x4B] = 0x20
	clientbound[V1_8][0x4C] = 0x1D

	serverbound[V1_8] = make(map[int]int)
	serverbound[V1_8][0x14] = 0x01
	serverbound[V1_8][0x01] = 0x02
	serverbound[V1_8][0x16] = 0x03
	serverbound[V1_8][0x15] = 0x04
	serverbound[V1_8][0x0F] = 0x05
	serverbound[V1_8][0x11] = 0x06
	serverbound[V1_8][0x0E] = 0x07
	serverbound[V1_8][0x0D] = 0x08
	serverbound[V1_8][0x17] = 0x09
	serverbound[V1_8][0x02] = 0x0A
	serverbound[V1_8][0x00] = 0x0B
	serverbound[V1_8][0x04] = 0x0C
	serverbound[V1_8][0x06] = 0x0D
	serverbound[V1_8][0x05] = 0x0E
	serverbound[V1_8][0x03] = 0x0F
	serverbound[V1_8][0x13] = 0x12
	serverbound[V1_8][0x07] = 0x13
	serverbound[V1_8][0x0B] = 0x14
	serverbound[V1_8][0x0C] = 0x15
	serverbound[V1_8][0x19] = 0x16
	serverbound[V1_8][0x09] = 0x17
	serverbound[V1_8][0x10] = 0x18
	serverbound[V1_8][0x12] = 0x19
	serverbound[V1_8][0x0A] = 0x1A
	serverbound[V1_8][0x18] = 0x1B
	serverbound[V1_8][0x08] = 0x1C

	// Hack 1.7.6
	clientbound[V1_7_6] = clientbound[V1_8]
	serverbound[V1_7_6] = serverbound[V1_8]

	// Hack 1.7.2
	clientbound[V1_7_2] = clientbound[V1_7_6]
	serverbound[V1_7_2] = serverbound[V1_7_6]

	// Hack 1.12
	clientbound[V1_12] = make(map[int]int)
	clientbound[V1_12][0x28] = 0x25
	for i := 0x25; i <= 0x27; i++ {
		clientbound[V1_12][i] = i + 1
	}
	for i := 0x30; i <= 0x34; i++ {
		clientbound[V1_12][i] = i + 1
	}
	for i := 0x35; i <= 0x49; i++ {
		clientbound[V1_12][i] = i + 2
	}
	for i := 0x4A; i <= 0x4B; i++ {
		clientbound[V1_12][i] = i + 3
	}

	serverbound[V1_12] = make(map[int]int)
	for i := 0x02; i <= 0x16; i++ {
		serverbound[V1_12][i] = i - 1
	}
	serverbound[V1_12][0x18] = 0x16
	for i := 0x1A; i <= 0x20; i++ {
		serverbound[V1_12][i] = i - 3
	}

	// Hack 1.12.1
	clientbound[V1_12_1] = copyHack(clientbound[V1_12])
	for i := 0x2B; i <= 0x4E; i++ {
		clientbound[V1_12_1][lastClientbound(V1_12, i)-1] = i
	}

	serverbound[V1_12_1] = copyHack(serverbound[V1_12])
	for i := 0x01; i <= 0x11; i++ {
		serverbound[V1_12_1][i] = serverbound[V1_12][i+1]
	}

	// Hack 1.12.2
	clientbound[V1_12_2] = clientbound[V1_12_1]
	serverbound[V1_12_2] = serverbound[V1_12_1]
}

func lastClientbound(proto Protocol, i int) int {
	for key, value := range clientbound[proto] {
		if value == i {
			return key
		}
	}
	return i
}

func copyHack(last map[int]int) map[int]int {
	newMap := make(map[int]int)
	for k, v := range last {
		newMap[k] = v
	}
	return newMap
}

func (player *Player) HackServerbound(id int) int {
	_, ok := serverbound[player.protocol]
	if ok {
		if val, ok := serverbound[player.protocol][id]; ok {
			return val
		} else {
			return id
		}
	}
	return id
}

func (player *Player) HackClientbound(id int) int {
	_, ok := clientbound[player.protocol]
	if ok {
		if val, ok := clientbound[player.protocol][id]; ok {
			return val
		} else {
			return id
		}
	}
	return id
}
