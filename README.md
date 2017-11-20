# TyphoonLimbo
### Lightweight minecraft limbo server

![1.10.X](https://img.shields.io/badge/1.10.X-ready-brightgreen.svg "1.10.X") ![1.11.X](https://img.shields.io/badge/1.11.X-ready-brightgreen.svg "1.11.X")
![1.12.X](https://img.shields.io/badge/1.12.X-ready-brightgreen.svg "1.12.X")
----
#### What is a limbo server ?
A limbo server is a connection keeper. You are only able to connect, nothing else

#### Minecraft protocol support

Minecraft compatible versions:

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
