package tcp

import (
	"fmt"
	"log"
	"net"
	"sync"
)

func Scanner(addr string) error {
	coon, err := net.Dial("tcp", addr)
	if err != nil {
		return err
	}
	defer coon.Close()
	return nil
}

func worker(id int, addrChan chan string, wg *sync.WaitGroup) {
	for addr := range addrChan {
		log.Printf("worker[%d] start, addr: %s\n", id, addr)

		err := Scanner(addr)
		if err != nil {
			log.Printf("[%s] connection fail, err: %s\n", addr, err)
			wg.Done()
			continue
		}

		log.Printf("worker[%d] done, addr: %s\n", id, addr)
		wg.Done()
	}
}

func workerMultiChan(id int, addr string, ports, results chan int) {
	for port := range ports {
		log.Printf("worker[%d] start, addr: %s:%d\n", id, addr, port)

		err := Scanner(fmt.Sprintf("%s:%d", addr, port))
		if err != nil {
			results <- 0
			log.Printf("worker[%d] [%s:%d] connection fail, err: %s\n", id, addr, port, err)
			continue
		}

		log.Printf("worker[%d] done, addr: %s:%d\n", id, addr, port)
		results <- port
	}
}
