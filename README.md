# TyphoonLimbo
## Lightweight minecraft limbo server

[![Build Status](https://travis-ci.org/TyphoonMC/TyphoonLimbo.svg?branch=master)](https://travis-ci.org/TyphoonMC/TyphoonLimbo)
----
### What is a limbo server ?
A limbo server is a fallback server able to handle a massive amount of simultaneous connections. The player spawns into the void then waits here. It can be used to keep players connected to a network after a lobby crashed or as an afk server.

### Minecraft protocol support

| Minecraft Version | Protocol Version | Supported | Comment                                             |
|-------------------|------------------|-----------|-----------------------------------------------------|
| 1.7.2 to 1.7.5    | 4                | true      | No boss bar, player list header/footer, compression |
| 1.7.6 to 1.7.10   | 5                | true      | No boss bar, player list header/footer              |
| 1.8 to 1.8.9      | 47               | true      | No boss bar                                         |
| 1.9               | 107              | true      |                                                     |
| 1.9.1             | 108              | true      |                                                     |
| 1.9.2             | 109              | true      |                                                     |
| 1.9.3 to 1.9.4    | 110              | true      |                                                     |
| 1.10 to 1.10.2    | 210              | true      |                                                     |
| 1.11              | 315              | true      |                                                     |
| 1.11.1 to 1.11.2  | 316              | true      |                                                     |
| 1.12              | 335              | true      |                                                     |
| 1.12.1            | 338              | true      |                                                     |
| 1.12.2            | 340              | true      |                                                     |

### How to build and start
```shell
go get github.com/TyphoonMC/go.uuid
go build
./TyphoonLimbo
```

### Security concerns
TyphoonLimbo may be used behind a Bungeecord proxy.

### Performance
#### Memory cost
Initial memory usage is about 580KB. You should consider an additonal 200KB per player connection (Used by socket buffers and packet wrappers)

#### CPU cost
Feel free to send me your metrics while using it.

My actual CPU cost is about 0% while keeping 10 clients. An older version of TyphoonLimbo was used on my network and handled about 800 connections correctly after a massive crash.
