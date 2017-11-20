package main

var (
	clientbound = make(map[Protocol]map[int]int)
  serverbound = make(map[Protocol]map[int]int)
	HACKS = [...]Protocol{V1_12,V1_12_1,V1_12_2}
)

func InitHacks() {
  // Hack 1.12
  clientbound[V1_12] = make(map[int]int);
  clientbound[V1_12][0x28] = 0x25
  for i := 0x25; i <= 0x27; i++ {
    clientbound[V1_12][i] = i+1
  }
  for i := 0x30; i <= 0x34; i++ {
    clientbound[V1_12][i] = i+1
  }
  for i := 0x35; i <= 0x49; i++ {
    clientbound[V1_12][i] = i+2
  }
  for i := 0x4A; i <= 0x4B; i++ {
    clientbound[V1_12][i] = i+3
  }

	serverbound[V1_12] = make(map[int]int);
  for i := 0x02; i <= 0x16; i++ {
    serverbound[V1_12][i] = i-1
  }
  serverbound[V1_12][0x18] = 0x16
  for i := 0x1A; i <= 0x20; i++ {
    serverbound[V1_12][i] = i-3
  }

  // Hack 1.12.1
	clientbound[V1_12_1] = copyHack(clientbound[V1_12]);
  for i := 0x2B; i <= 0x4E; i++ {
    clientbound[V1_12_1][i] = lastClientbound(V1_12, i+1)
  }

  serverbound[V1_12_1] = copyHack(serverbound[V1_12]);
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
  for k,v := range last {
    newMap[k] = v
  }
  return newMap
}

func (player *Player) HackServerbound(id int) int {
	for _, x := range HACKS {
		if x == player.protocol {
			if val, ok := serverbound[x][id]; ok {
				return val
			} else {
				return id
			}
		}
	}
	return id
}

func (player *Player) HackClientbound(id int) int {
	for _, x := range HACKS {
		if x == player.protocol {
			if val, ok := clientbound[x][id]; ok {
				return val
			} else {
				return id
			}
		}
	}
	return id
}
