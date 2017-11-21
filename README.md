# TyphoonLimbo
### Lightweight minecraft limbo server

![1.7.X](https://img.shields.io/badge/1.7.X-partial-orange.svg "1.7.X partial")
![1.8.X](https://img.shields.io/badge/1.8.X-ready-brightgreen.svg "1.8.X")
![1.9.X](https://img.shields.io/badge/1.9.X-ready-brightgreen.svg "1.9.X")
![1.10.X](https://img.shields.io/badge/1.10.X-ready-brightgreen.svg "1.10.X") ![1.11.X](https://img.shields.io/badge/1.11.X-ready-brightgreen.svg "1.11.X")
![1.12.X](https://img.shields.io/badge/1.12.X-ready-brightgreen.svg "1.12.X")
----
#### What is a limbo server ?
A limbo server is a fallback server able to handle a massive amount of simultaneous connections. The player spawns into the void then waits here. It can be used to keep players connected to a network after a lobby crashed or as an afk server.

#### Minecraft protocol support

Minecraft compatible versions:

* 1.7.2 to 1.7.5 (4)
* 1.7.6 to 1.7.10 (5)
* 1.8 to 1.8.9 (47)
* 1.9 (107)
* 1.9.1 (108)
* 1.9.2 (109)
* 1.9.3, 1.9.4 (110)
* 1.10, 1.10.1, 1.10.2 (210)
* 1.11 (315)
* 1.11.1, 1.11.2 (316)
* 1.12 (335)
* 1.12.1 (338)
* 1.12.2 (340)

#### How to build and start
```shell
go get github.com/satori/go.uuid
go build
./TyphoonLimbo
```

#### Security concerns
TyphoonLimbo may actually be used behind a Bungeecord proxy.
