groxi
=============
[Gorox](https://github.com/klustic/gorocks) Socks Proxy Improved.

This is a result of me wanting to learn more golang and how gorox works. I really like gorox, but cert generation seemed like a waste of time. I also wanted to have the ability to use a single binary most of the time, but still be able to have a standalone relay or server if the need arises. Maybe someone else will find this useful.

`groxi`
------------
```
Just drop the binary and run it. Certificates are auto generated

Usage:
  groxi [command]

Available Commands:
  help        Help about any command
  relay       Connects to a server and allows traffic to pass through it
  server      Server listen for socks clients and relay clients

Flags:
  -d, --debug   Debug mode
  -h, --help    help for groxi

Use "groxi [command] --help" for more information about a command.
```

### `relay`

```
Connects to a server and allows traffic to pass through it. For example:

To fail after one connection attempt:
    groxi relay -f 1 -t 10.0.1.10:8081


NOTE: exponential falloff on retries. sleepMilis(failedAttemptCount * failedAttemptCount * 100)
bash snippet to help you decide what number to use.
    last=0;for i in {1..30}; do me=$(echo "$last+($i^2*100)"|bc); echo "$me/1000/60"|bc; last=$me  ;done

Usage:
  groxi relay [flags]

Flags:
  -f, --fail int        The number of connections to try before giving up. 30 is about 15 minutes. (default 30)
  -h, --help            help for relay
  -t, --tunnel string   The bind address on which to accept tunnel connections (default "127.0.0.1:8081")

Global Flags:
  -d, --debug   Debug mode
```

### `server`

```
Setup 2 listeners one for socks clients, one for relay clients.

Usage:
  groxi server [flags]

Flags:
  -h, --help            help for server
  -s, --socks string    The bind address on which to accept SOCKSv5 clients (default "127.0.0.1:1080")
  -t, --tunnel string   The bind address on which to accept tunnel connections (default "0.0.0.0:8081")

Global Flags:
  -d, --debug   Debug mode
```

*******

`groxi_relay`
-------------------

```
Usage of groxi_relay:
  -f int
      The number of connections to try before giving up. NOTE: Exponential fall off (default 30)
  -t string
      Address to connect to server on. (default "127.0.0.1:8081")
  -v	Print groxi version
```

******

`groxi_server`
-------------------

```
Usage of groxi_server:
  -s string
      Address to accept socks connections on. (default "127.0.0.1:1080")
  -t string
      Address to accept relay connections on. (default "127.0.0.1:8081")
  -v	prints groxi version
```

**************

Building
--------

```bash
make
```

Simply run `make` to build binaries for linux, windows, and darwin. This will compile everything into `dist/`.

```
├── darwin
│   └── amd64
│       └── bin
│           ├── grx
│           ├── grx_relay
│           └── grx_server
├── linux
│   └── amd64
│       └── bin
│           ├── grx
│           ├── grx_relay
│           └── grx_server
└── windows
    └── amd64
        └── bin
            ├── grx
            ├── grx_relay
            └── grx_server
```


### UPX

If you have `upx` installed, you can run `make upx` to compress all the things and put them in a directory named `upx/` next to the bin directory.
