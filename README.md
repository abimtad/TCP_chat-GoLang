# TCP Chat (Go)

A small TCP chat server built in Go. Clients connect over plain TCP and use slash commands to set nicknames, join rooms, list rooms, send messages, and quit.

## Features

- Multiple rooms with join/leave notifications
- Per-connection nicknames
- Simple command parsing over TCP
- Room broadcast messaging

## Commands

- /nick NAME
- /join ROOM_NAME
- /rooms
- /msg MESSAGE
- /quit

## Run

From the repo root:

```
go run ./chat
```

The server listens on port 8888.

## Try It

In separate terminals:

```
# terminal 1
nc localhost 8888

# terminal 2
nc localhost 8888
```

Then use commands like:

```
/nick abel
/join general
/msg hello everyone
```

## Project Layout

- chat/main.go: TCP listener and connection loop
- chat/server.go: server, command dispatch, room logic
- chat/client.go: client read loop and command parsing
- chat/room.go: room member list and broadcast
- chat/command.go: command IDs and struct

## Notes

This is a minimal learning project and uses plain TCP (no TLS) and a simple line-based protocol.
