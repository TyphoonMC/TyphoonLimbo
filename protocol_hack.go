package main

var (
	hackMap = make(map[Protocol]map[int]int)
	HACKS = [...]Protocol{V1_12}

	hackMap1_12 = make(map[int]int)
)

func InitHacks() {
	hackMap[V1_12] = hackMap1_12;

	hackMap1_12[0x02] = 0x01
	hackMap1_12[0x03] = 0x02
	hackMap1_12[0x0C] = 0x0B
}

func (player *Player) Hack(id int) int {
	for _, x := range HACKS {
		if x == player.protocol {
			if val, ok := hackMap[x][id]; ok {
				return val
			} else {
				return id
			}
		}
	}
	return id
}
