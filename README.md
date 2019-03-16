# TyphoonLimbo
## Lightweight minecraft limbo server

[![Build Status](https://travis-ci.org/TyphoonMC/TyphoonLimbo.svg?branch=master)](https://travis-ci.org/TyphoonMC/TyphoonLimbo)
![stability-stable](https://img.shields.io/badge/stability-stable-green.svg)
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
| 1.13              | 393              | true      |                                                     |
| 1.13.1            | 401              | true      |                                                     |
| 1.13.2            | 404              | true      |                                                     |

#### Snapshot support
TyphoonLimbo is able to load [TyphoonDep protocol-map modules](https://github.com/TyphoonMC/TyphoonDep/tree/master/protocol-map) to add a partial snapshots support.

All json files are loaded from the "modules" folder in the same directory as the TyphoonLimbo binary.

You can generate a protocol-map module with the [fetcher](https://github.com/TyphoonMC/TyphoonDep/tree/master/protocol-map/fetcher). Just pick the "page" link of your wanted version on [wiki.vg](http://wiki.vg/Protocol_version_numbers) as argument of the fecther and the magic will happen.

### How to build and start
```shell
git clone https://github.com/TyphoonMC/TyphoonLimbo.git
cd TyphoonLimbo
go build
./TyphoonLimbo
```
