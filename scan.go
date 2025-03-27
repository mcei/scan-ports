// A tiny console application to scan system ports
package main

import (
	"fmt"
	"net"
	"os"
	"sync"
	"time"
)

const (
	maxRangePorts = 1024
	usage         = `Example usage:
	./scan 192.168.0.1
	./scan localhost
	./scan scanme.nmap.org
	`
)

var dialer = &net.Dialer{
	Timeout:   1 * time.Second,
	KeepAlive: 1 * time.Second,
}

// Scan port on a host address specified
func Scan(host string, port int, wg *sync.WaitGroup) {
	defer wg.Done()
	address := net.JoinHostPort(host, fmt.Sprintf("%d", port))
	connection, err := dialer.Dial("tcp", address)
	if err != nil {
		return
	}
	defer connection.Close()
	fmt.Println(connection.RemoteAddr().String(), "is open")
}

// Validate IP or hostname using local resolver
func Validate(host string) bool {
	// First validate string as an IP
	if _, err := net.LookupIP(host); err == nil {
		return true
	}
	// Then validate string as a hostname
	if _, err := net.LookupHost(host); err == nil {
		return true
	}
	return false
}

func main() {
	args := os.Args
	if len(args) <= 1 {
		fmt.Println(usage)
		os.Exit(1)
	}

	address := args[1]
	if ok := Validate(address); !ok {
		fmt.Println(usage)
		fmt.Println("Invalid host address:", address)
		os.Exit(1)
	}

	var wg sync.WaitGroup
	for port := 0; port < maxRangePorts; port++ {
		wg.Add(1)
		go Scan(address, port, &wg)
	}
	wg.Wait()
}
