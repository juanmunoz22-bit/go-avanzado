package main

import (
	"flag"
	"fmt"
	"net"
	"sync"
)

var site = flag.String("website", "scanme.nmap.org", "Port scan in a given URL")

func main() {
	//1, 2, 3, 4..., 99..., x
	//sitio:1, sitio:2..., sitio:99..., x
	// 1 -> Open..., 99 -> Closed...
	flag.Parse()
	var wg sync.WaitGroup
	for i := 1; i < 65535; i++ {
		wg.Add(1)
		go func(port int) {
			defer wg.Done()
			conn, err := net.Dial("tcp", fmt.Sprintf("%s:%d", *site, port))
			if err != nil {
				return
			}
			conn.Close()
			fmt.Printf("Port %d is open", port)
		}(i)
	}
	wg.Wait()
}
