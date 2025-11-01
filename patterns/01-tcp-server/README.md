# TCP Server Pattern

Simple TCP Echo Server demonstrating low-level network programming with Go's `net` package.

## Concepts

- **TCP Basics**: Connection-oriented, reliable, bidirectional communication
- **Goroutines**: Concurrent client handling (one goroutine per connection)
- **net.Listen()**: Opens TCP port and listens for connections
- **net.Accept()**: Blocks until client connects
- **bufio.Scanner**: Stream-based reading from connection

## Usage

```bash
# Start server
go run main.go

# Connect with telnet (in another terminal)
telnet localhost 8080

# Type messages - server echoes in UPPERCASE
> hello
Echo: HELLO

> quit
Connection closed by foreign host.
```

## Architecture

```
TCP Server (:8080)
│
├─ Client 1 → Goroutine 1
├─ Client 2 → Goroutine 2
└─ Client 3 → Goroutine 3
```

## Why TCP?

- **Foundation for HTTP**: HTTP runs on top of TCP
- **Reliable**: Guaranteed packet delivery
- **Order Preservation**: Messages arrive in sequence
- **Flow Control**: Prevents overwhelming receiver

## Next Steps

See `02-http-server` for building on top of TCP with HTTP protocol.
