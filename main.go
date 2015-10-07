package main

import (
	"flag"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net"
	"os"
)

func main() {
	var (
		flagSet = flag.NewFlagSet("sink", flag.ExitOnError)
		mode    = flagSet.String("mode", "udp", "listen on `udp` or `tcp`")
		port    = flagSet.Int("port", 0, "port to listen on")
	)
	if err := flagSet.Parse(os.Args[1:]); err != nil {
		log.Fatal(err)
	}
	if *mode != "udp" && *mode != "tcp" {
		log.Fatal("invalid mode: expect `tcp` or `udp`")
	}
	if *mode == "udp" {
		addr, err := net.ResolveUDPAddr(*mode, fmt.Sprintf(":%d", *port))
		if err != nil {
			log.Fatal(err)
		}
		ln, err := net.ListenUDP("udp", addr)
		if err != nil {
			log.Fatal(err)
		}
		io.Copy(ioutil.Discard, ln)
	} else {
		ln, err := net.Listen(*mode, fmt.Sprintf(":%d", *port))
		if err != nil {
			log.Fatal(err)
		}
		for {
			c, err := ln.Accept()
			if err != nil {
				log.Fatal(err)
			}
			go io.Copy(ioutil.Discard, c)
		}
	}
}
