package main

import (
	"bufio"
	"bytes"
	"fmt"
	"io"
	"net"
	"time"

	"github.com/joepaul-martin/image-server/imageProcessor"
)

func main() {
	p := make([]byte, 2048)
	conn, err := net.Dial("udp", "127.0.0.1:80")
	if err != nil {
		fmt.Printf("Some error %v", err)
		return
	}

	fileName := "./image.jpg"

	data := imageProcessor.ImageToBytes(fileName)

	writeLargeBytes(conn, data)
	_, err = bufio.NewReader(conn).Read(p)
	if err == nil {
		fmt.Printf("%s\n", p)
	} else {
		fmt.Printf("Some error %v\n", err)
	}
	conn.Close()
}

func writeLargeBytes(conn net.Conn, data []byte) {
	reader := bytes.NewReader(data)
	fmt.Println(len(data))
	fmt.Printf("Size of reader: %v\n", reader.Size())
	for {
		buf := make([]byte, 2048)
		_, err := reader.Read(buf)
		if err == io.EOF {
			fmt.Println("Completed reading the data")
			break
		}
		_, err = conn.Write(buf)
		if err != nil {
			fmt.Printf("Some error %v", err)
			return
		}
		time.Sleep(time.Millisecond * 10)
	}
	_, err := conn.Write([]byte("END"))
	if err != nil {
		fmt.Printf("Some error %v", err)
		return
	}
}
