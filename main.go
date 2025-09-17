package main

import (
	"flag"
	"net"
	"strconv"
	"strings"
	"time"
)

func main() {
	target := flag.String("target", "localhost", "Target host to scan")
	scanRange := flag.String("range", "1-1024", "Port range to scan (e.g., 1-1024)")
	flag.Parse()

	if err := checkForInvalidHostname(*target); err != nil {
		println("Invalid hostname:", err.Error())
		return
	}

	portRangeArr := strings.Split(*scanRange, "-")
	startPort, err1 := strconv.Atoi(portRangeArr[0])
	endPort, err2 := strconv.Atoi(portRangeArr[1])

	if err1 != nil || err2 != nil || startPort < 1 || endPort > 65535 || startPort > endPort {
		println("Invalid port range. Please use the format start-end (e.g., 1-1024).")
		return
	}

	scanPorts(*target, startPort, endPort)
}

func scanPorts(target string, startPort int, endPort int) {

	for port := startPort; port <= endPort; port++ {
		address := net.JoinHostPort(target, strconv.Itoa(port))
		conn, err := net.DialTimeout("tcp", address, time.Second)
		if err != nil {
			continue
		}
		conn.Close()
		println("Port", port, "is open")
	}
}

func checkForInvalidHostname(hostname string) error {
	_, err := net.LookupHost(hostname)
	return err
}
