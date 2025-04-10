package tcp

import (
	"fmt"
	"log"
	"net"
	"testing"
	"time"
)

func TestEcho(t *testing.T) {
	listener, err := net.Listen("tcp", ":20080")
	if err != nil {
		log.Fatalln("unable to bind to port.")
	}
	log.Println("listening on 0.0.0.0:20080.")

	for {
		conn, err := listener.Accept()
		log.Println("received connection.")
		if err != nil {
			log.Fatalln("unable to accept connection.")
		}
		connId := fmt.Sprintf("conn-%d", time.Now().Unix())
		//go echo(conn, connId)
		go echoWithBufIO(conn, connId)
		//go echoWithIOCopy(conn)
	}
}

func TestClient(t *testing.T) {
	conn, err := net.Dial("tcp", ":20080")
	if err != nil {
		log.Fatalln("unable to connect to port.")
	}

	_, err = conn.Write([]byte("hello world."))
	if err != nil {
		log.Fatalln("unable to write data, err: ", err)
	}

	echo(conn, "client")
}
