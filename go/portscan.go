package main
// Run with:
// go run portscan.go www.google.com


import (
	"strconv"
	"log"
	"net"
	"os"
	"time"
)

func printUsage() {
	log.Println("Usage: ")
	log.Println("	go run portscan.go <host> ")
	log.Println("Example: ")
	log.Println("	go run portscan.go www.google.com ")
	log.Println("	go run portscan.go 8.8.8.8 ")
}

func TestTcpConnect(host string, port int, doneChannel chan bool) {
	timeoutLength := 5 * time.Second
	conn, err := net.DialTimeout("tcp", host +":" + strconv.Itoa(port), timeoutLength)
	if err != nil {
		doneChannel <- false
		return
	}
	conn.Close()
	log.Printf("[+] %d connected", port)
	doneChannel <- true
}

func main() {
	if len(os.Args) == 1 {
		log.Println("No arguments.")
		printUsage()
		os.Exit(1)
	}
	doneChannel := make(chan bool)
	activeThreadCount := 0
	log.Println("Scanning host: " + os.Args[1])
	for PortNumber := 1; PortNumber <= 65535; PortNumber++ {
		activeThreadCount++
		go TestTcpConnect(os.Args[1], PortNumber, doneChannel)
	}		
	for {
		<- doneChannel
		activeThreadCount--;
		if activeThreadCount == 0 {
			break
		}
	}
}