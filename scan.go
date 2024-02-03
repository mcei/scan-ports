package main

import (
	"fmt"
	"net"
	"os"
	"sync"
	"time"
)

const (
	Timeout       = 1 * time.Second
	MaxRangePorts = 1024
)

// scan port on a host address specified
func scan(host string, port int, wg *sync.WaitGroup) {
	defer wg.Done()

	address := fmt.Sprintf("%s:%d", host, port)

	connection, err := net.DialTimeout("tcp", address, Timeout)
	if err != nil {
		return
	}

	defer connection.Close()

	fmt.Println(connection.RemoteAddr().String(), "is open")
}

// validate IP or host address
func validate(host string) bool {
	if _, err := net.LookupIP(host); err == nil {
		return true
	}

	if _, err := net.LookupHost(host); err == nil {
		return true
	}

	return false
}

func main() {
	usage := `Example usage:
	./scan 192.168.0.1
	./scan scanme.nmap.org
	`
	args := os.Args
	if len(args) <= 1 {
		fmt.Println(usage)
		fmt.Println("Address expected")

		os.Exit(1)
	}

	address := args[1]

	ok := validate(address)
	if !ok {
		fmt.Println(usage)
		fmt.Println("Invalid host address:", address)

		os.Exit(1)
	}

	var wg sync.WaitGroup

	for port := 0; port < MaxRangePorts; port++ {
		wg.Add(1)
		go scan(address, port, &wg)
	}

	wg.Wait()
}
