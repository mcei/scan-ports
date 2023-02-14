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

func scan(host string, port int, wg *sync.WaitGroup) {
	defer wg.Done()

	address := fmt.Sprintf("%s:%d", host, port)

	connection, err := net.DialTimeout("tcp", address, Timeout)
	if err != nil {
		return
	}

	fmt.Println(connection.RemoteAddr().String(), "is open")

	connection.Close()
}

func validate(host string) bool {

	// Verify if IP or host address is valid

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

	for port := range make([]int, MaxRangePorts) {
		wg.Add(1)
		go scan(address, port, &wg)
	}

	wg.Wait()
}
