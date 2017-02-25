package main

import (
	"io"
	"log"
	"net"
	"runtime"

	"github.com/mdlayher/arp"
)

func main() {
	is, err := net.Interfaces()
	if err != nil {
		log.Println(err)
		return
	}
	for _, v := range is {
		go func(v net.Interface) {
			c, err := arp.NewClient(&v)
			if err != nil {
				log.Println(err)
				return
			}
			for {
				pkt, eth, err := c.Read()
				if err != nil {
					if err == io.EOF {
						log.Println("EOF")
						break
					}
					log.Println(err)
					return
				}
				if pkt.Operation == arp.OperationRequest {
					log.Println(pkt, eth)
				}
			}
		}(v)
	}
	runtime.Goexit()
}
