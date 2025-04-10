package tcp

import (
	"io"
	"log"
	"net"
	"testing"
)

func TestFoo(t *testing.T) {
	var (
		reader FooReader
		writer FooWriter
	)

	//input := make([]byte, 4096)
	//n, err := reader.Read(input)
	//if err != nil {
	//	log.Fatalln("unable to read data, err: ", err)
	//}
	//fmt.Printf("read %d bytes from stdin\n", n)
	//
	//n, err = writer.Write(input)
	//if err != nil {
	//	log.Fatalln("unable to write data.")
	//}
	//fmt.Printf("write %d bytes from stdout\n", n)

	_, err := io.Copy(&writer, &reader)
	if err != nil {
		log.Fatalln("unable to read/write data.")
	}
}

func TestProxy(t *testing.T) {
	listener, err := net.Listen("tcp", ":8080")
	if err != nil {
		log.Fatalln("unable to bind to port.")
	}
	log.Println("listening on 0.0.0.0:8080.")

	for {
		conn, err := listener.Accept()
		log.Println("received connection.")
		if err != nil {
			log.Fatalln("unable to accept connection.")
		}
		go handler(conn)
	}
}
