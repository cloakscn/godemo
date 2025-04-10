package tcp

import (
	"fmt"
	"io"
	"log"
	"net"
	"os"
)

type FooReader struct{}

func (fooReader *FooReader) Read(b []byte) (int, error) {
	fmt.Print("in > ")
	return os.Stdin.Read(b)
}

type FooWriter struct{}

func (fooWriter *FooWriter) Write(b []byte) (int, error) {
	fmt.Print("out> ")
	return os.Stdout.Write(b)
}

func handler(src net.Conn) {
	dst, err := net.Dial("tcp", ":20080")
	if err != nil {
		log.Fatalln("Unable to connect to our unreachable host")
	}
	defer dst.Close()

	go func() {
		if _, err = io.Copy(dst, src); err != nil {
			log.Fatalln(err)
		}
	}()

	if _, err = io.Copy(src, dst); err != nil {
		log.Fatalln(err)
	}
}
