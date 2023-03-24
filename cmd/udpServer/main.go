package main

import (
	"fmt"
	"net"

	"github.com/joepaul-martin/image-server/imageProcessor"
)

func main() {
	addr := net.UDPAddr{
		IP:   net.ParseIP("127.0.0.1"),
		Port: 80,
	}

	ser, err := net.ListenUDP("udp", &addr)
	if err != nil {
		fmt.Printf("There was some error, error : %v\n", err)
	}
	var buf []byte
	count := 0
	for {
		p := make([]byte, 2048)
		_, remoteaddr, err := ser.ReadFromUDP(p)
		fmt.Printf("Read a message from %v, count : %v\n", remoteaddr, count)
		if err != nil {
			fmt.Printf("Some error  %v", err)
			continue
		}
		if string(p) != "END" && count == 200 {
			buf = append(buf, p...)
			fmt.Printf("Adding to the buffer, count: %v\n", count)
		}
		if count == 412 {
			fmt.Printf("Size of the buffer : %v", len(buf))
			fmt.Printf("Reached the end, last message: %s\n, starting image processing\n", p)
			err = imageProcessor.BytesToImage(buf, "./outImage.jpeg")
			if err != nil {
				fmt.Printf("Error while saving the bytes to image, error: %v\n", err)
			} else {
				go sendResponse(ser, remoteaddr)
			}

		}
		count++
	}
}

func sendResponse(conn *net.UDPConn, addr *net.UDPAddr) {
	_, err := conn.WriteToUDP([]byte("From server: Hello I got your image "), addr)
	if err != nil {
		fmt.Printf("Couldn't send response %v", err)
	}
}
