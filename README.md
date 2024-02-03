[![Go](https://github.com/mcei/scan-ports/actions/workflows/go.yml/badge.svg)](https://github.com/mcei/scan-ports/actions/workflows/go.yml) [![GoDoc](https://godoc.org/github.com/mcei/scan-ports?status.svg)](https://godoc.org/github.com/mcei/scan-ports) [![Go Report Card](https://goreportcard.com/badge/github.com/mcei/scan-ports)](https://goreportcard.com/report/github.com/mcei/scan-ports)

The program is looking over open system ports of range 0 - 1023 on a host specified

Build:

`go build scan.go`

Usage: 

`./scan 192.168.0.1`

`./scan localhost`

`./scan scanme.nmap.org`
