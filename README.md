# TyphoonLimbo
### Lightweight minecraft limbo server

![1.10.X](https://img.shields.io/badge/1.10.X-ready-brightgreen.svg "1.10.X") ![1.11.X](https://img.shields.io/badge/1.11.X-ready-brightgreen.svg "1.11.X")
![1.12](https://img.shields.io/badge/1.12-partial-orange.svg "1.12")
----
#### What is a limbo server ?
A limbo server is a connection keeper. You are only able to connect, nothing else

#### Minecraft protocol support

Confirmed:

* 1.10, 1.10.1, 1.10.2 (210)
* 1.11 (315)
* 1.11.1, 1.11.2 (316)
* 1.12 (335)

Some other protocol versions may work

#### How to build and start
```shell
go get github.com/satori/go.uuid
go build
./TyphoonLimbo
```

#### Security concerns
TyphoonLimbo may actually be used behind a Bungeecord proxy.
