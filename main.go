package main

import (
	"flag"
	"fmt"
	"net"
	"os"
	"strconv"
	"strings"
	"time"
)

func main() {
	target := flag.String("target", "localhost", "Target host to scan")
	scanRange := flag.String("range", "1-1024", "Port range to scan (e.g., 1-1024)")
	help := flag.Bool("help", false, "Display help")

	for i, arg := range os.Args {
		switch arg {
		case "-t", "--target":
			os.Args[i] = "-target"
		case "-r", "--range":
			os.Args[i] = "-range"
		case "-h", "--help":
			os.Args[i] = "-help"
		}
	}

	flag.Parse()

	if *help {
		printHelp()
		return
	}

	hosts := strings.Split(*target, ",")
	for _, host := range hosts {

		if err := checkForInvalidHostname(host); err != nil {
			fmt.Printf("Invalid hostname '%s': %s\n", host, err.Error())
			return
		}

		startPort, endPort, err := parsePortRange(*scanRange)
		if err != nil {
			return
		}

		fmt.Printf("\nScanning host: %s\n\n", host)
		scanPorts(host, startPort, endPort)
	}

}

func scanPorts(target string, startPort int, endPort int) {
	openPortFound := false

	for port := startPort; port <= endPort; port++ {
		address := net.JoinHostPort(target, strconv.Itoa(port))
		conn, err := net.DialTimeout("tcp", address, time.Second)
		if err != nil {
			continue
		}
		conn.Close()
		println("Port", port, "is open")
		openPortFound = true
	}

	if !openPortFound {
		println("No open ports found in the specified range.\n")
	}
}

func parsePortRange(scanRange string) (int, int, error) {
	portRangeArr := strings.Split(scanRange, "-")
	startPort, err1 := strconv.Atoi(portRangeArr[0])
	endPort, err2 := strconv.Atoi(portRangeArr[1])

	if err1 != nil || err2 != nil || startPort < 1 || endPort > 65535 || startPort > endPort {
		println("Invalid port range. Please use the format start-end (e.g., 1-1024).")
		return 0, 0, fmt.Errorf("invalid port range")
	}

	return startPort, endPort, nil
}

func checkForInvalidHostname(hostname string) error {
	_, err := net.LookupHost(hostname)
	return err
}

func printHelp() {
	helpText := `
Usage: portscanner [options]

Options:
  -t, --target   Target host to scan (default: localhost). You can specify multiple hosts separated by commas.
  -r, --range    Port range to scan in the format start-end (default: 1-1024).
  -h, --help     Display this help message.

Example:
  portscanner -t example.com,localhost -r 20-80
`
	fmt.Println(helpText)
}
