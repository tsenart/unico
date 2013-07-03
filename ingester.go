package main

import (
	"log"
	"net"
)

const BUFFER_SIZE = 1024

type Ingester struct {
	listener *net.UDPConn
	visits   chan Visit
}

func NewIngester(addr string, visits chan Visit) (*Ingester, error) {
	in := &Ingester{visits: visits}

	udpAddr, err := net.ResolveUDPAddr("udp4", addr)
	if err != nil {
		return nil, err
	}
	ln, err := net.ListenUDP("udp4", udpAddr)
	if err != nil {
		return nil, err
	}
	in.listener = ln

	return in, nil
}

func (in *Ingester) Start() {
	log.Printf("Ingester: started")
	go in.loop()
}

func (in *Ingester) Close() error {
	if err := in.listener.Close(); err != nil {
		log.Printf("Ingester Close: %s", err)
		return err
	}
	close(in.visits)
	log.Printf("Ingester: done")
	return nil
}

func (in *Ingester) loop() {
	for {
		in.ingest()
	}
}

func (in *Ingester) ingest() error {
	buffer := make([]byte, BUFFER_SIZE)

	n, addr, err := in.listener.ReadFromUDP(buffer)
	if err != nil {
		log.Printf("Ingester: error %s (continuing)", err)
		return err
	}
	log.Printf("Ingester: %d %s %s", n, addr, buffer)

	if visit, err := ParseVisit(buffer[0:n]); err != nil {
		log.Printf("Ingester: %s (continuing)", err)
		return err
	} else {
		in.visits <- visit
	}
	return nil
}
