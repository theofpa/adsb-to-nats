package main

import (
	"bufio"
	"fmt"
	nats "github.com/nats-io/nats.go"
	"net"
	"os"
)

func main() {
	// Get the ADS-B server and NATS server from environment variables
	adsbServer := os.Getenv("ADS_B_SERVER")
	natsServer := os.Getenv("NATS_SERVER")

	// Connect to the NATS server
	nc, err := nats.Connect(natsServer)
	if err != nil {
		panic(err)
	}
	defer nc.Close()

	// Connect to the ADS-B server
	conn, err := net.Dial("tcp", adsbServer+":30003")
	if err != nil {
		panic(err)
	}
	defer conn.Close()

	// Read messages from the ADS-B server and publish them to NATS
	scanner := bufio.NewScanner(conn)
	for scanner.Scan() {
		msg := scanner.Text()
		fmt.Println(msg) // For debugging purposes only
		nc.Publish("adsb", []byte(msg))
	}
}
