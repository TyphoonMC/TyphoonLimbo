# DragonLimbo
## Lightweight minecraft limbo server

[![Build Status](https://travis-ci.org/TyphoonMC/TyphoonLimbo.svg?branch=master)](https://travis-ci.org/TyphoonMC/TyphoonLimbo)
![stability-wip](https://img.shields.io/badge/stability-work_in_progress-lightgrey.svg)
----
### What is a limbo server ?
A limbo server is a fallback server able to handle a massive amount of simultaneous connections. The player spawns into the void then waits here. It can be used to keep players connected to a network after a lobby crashed or as an afk server.

### DragonLimbo vs TyphoonLimbo
DragonLimbo aims to provide the same abilities as TyphoonLimbo to the Bedrock Engine based games

### Bedrock protocol support

| Bedrock Version   | Protocol Version | Supported | Comment |
|-------------------|------------------|-----------|---------|
| 1.2.10            | 200              | Will be   |         |

### How to build and start
```shell
go get github.com/TyphoonMC/go.uuid
go build
./DragonLimbo
```