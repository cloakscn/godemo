package tcp

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"net"
)

func echo(conn net.Conn, connId string) {
	defer conn.Close()

	b := make([]byte, 1024)
	for {
		n, err := conn.Read(b[0:])
		if err == io.EOF {
			log.Println("client disconnected.")
			break
		}
		if err != nil {
			log.Println("unexpected error.")
			break
		}
		log.Printf("received %d bytes: %s\n", n, string(b))

		log.Println("writing data.")
		msg := fmt.Sprintf("[%s] send: %s", connId, b[0:n])
		if _, err = conn.Write([]byte(msg)); err != nil {
			log.Fatalln("unable to write data.")
		}
	}
}

func echoWithIOCopy(conn net.Conn) {
	defer conn.Close()

	if _, err := io.Copy(conn, conn); err != nil {
		log.Fatalln("unable to copy conn, err: ", err)
	}
}

func echoWithBufIO(conn net.Conn, connId string) {
	defer conn.Close()

	reader := bufio.NewReader(conn)
	writer := bufio.NewWriter(conn)

	for {
		s, err := reader.ReadString('\n')
		if err != nil {
			log.Println("> unable to read data.")
		}
		log.Printf("> [%s] received %d bytes: %s\n", connId, len(s), s)

		msg := fmt.Sprintf("> [%s] send: %s", connId, s)
		if _, err := writer.WriteString(msg); err != nil {
			log.Fatalln("> unable to write data.")
		}

		writer.Flush()
	}
}
