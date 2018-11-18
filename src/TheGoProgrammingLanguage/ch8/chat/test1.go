// Description: abc

// Copyright (C) 2018 liyunteng
// Last-Updated: <2018/11/18 05:16:28 liyunteng>

package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"strconv"
	"time"
)

type client struct {
	id   int
	name string
	addr net.Addr

	ch      chan string
	conn    net.Conn
	timeout time.Duration
	toch    chan struct{}
}

var (
	listenAddr = "localhost:8080"
	entering   = make(chan client)
	leaving    = make(chan client)
	messages   = make(chan string)
	count      = 0
	timeout    = 10 * time.Second
)

func broadcast() {
	clients := make(map[client]bool)

	for {
		select {
		case c := <-entering:
			clients[c] = true

		case msg := <-messages:
			for c := range clients {
				c.ch <- msg
			}

		case c := <-leaving:
			delete(clients, c)
			close(c.ch)
		}
	}
}

func (c *client) clientWriter() {
	for msg := range c.ch {
		fmt.Fprintln(c.conn, msg)
	}
}

func (c *client) clientReader() {
	input := bufio.NewScanner(c.conn)
	for input.Scan() {
		messages <- c.name + ": " + input.Text()
		c.toch <- struct{}{}
	}
}

func handleConn(conn net.Conn) {
	c := client{
		id:      count,
		name:    strconv.Itoa(count),
		addr:    conn.RemoteAddr(),
		ch:      make(chan string),
		conn:    conn,
		timeout: timeout,
		toch:    make(chan struct{}),
	}
	count++
	go c.clientWriter()
	go c.clientReader()

	c.ch <- "You are " + c.name
	messages <- c.name + "(" + c.addr.String() + ")" + " has arrived"
	entering <- c

loop:
	for {
		select {
		case <-time.After(c.timeout):
			break loop
		case <-c.toch:

		}
	}
	leaving <- c
	messages <- c.name + "(" + c.addr.String() + ")" + " has left"
	conn.Close()
}

func main() {
	listener, err := net.Listen("tcp", listenAddr)
	if err != nil {
		log.Fatal(err)
	}

	go broadcast()

	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Print(err)
			continue
		}
		go handleConn(conn)
	}
}
