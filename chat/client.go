package main

import (
	"bufio"
	"net"
)

type client struct {
	conn     net.Conn
	nick     string
	room     *room
	commands chan<- command
}

func (c *client) readInput() {
	reader := bufio.NewReader(c.conn)
	for {
		_, err := reader.ReadString('\n')
		if err != nil {
			return
		}
	}
}
