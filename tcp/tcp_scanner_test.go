package tcp

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"sort"
	"strconv"
	"sync"
	"testing"
)

const (
	addr = "itycu.cn"
	nums = 65535
)

// Tcp Scanner
func TestScanner(t *testing.T) {
	for i := 1; i <= nums; i++ {
		err := Scanner(fmt.Sprintf("%s:%s", addr, strconv.Itoa(i)))
		assert.Nil(t, err)
	}
}

func TestConcurrentScanner(t *testing.T) {
	wg := sync.WaitGroup{}
	for i := 1; i <= nums; i++ {
		wg.Add(1)
		go func(i int) {
			defer wg.Done()
			err := Scanner(fmt.Sprintf("%s:%s", addr, strconv.Itoa(i)))
			assert.Nil(t, err)
		}(i)
	}
	wg.Wait()
}

func TestConcurrentScannerWithWorker(t *testing.T) {
	addrChan := make(chan string, 100)
	defer close(addrChan)

	var wg sync.WaitGroup
	for i := 0; i < cap(addrChan); i++ {
		// consumer
		go worker(i, addrChan, &wg)
	}

	for i := 1; i <= nums; i++ {
		// provider
		wg.Add(1)
		addrChan <- fmt.Sprintf("%s:%s", addr, strconv.Itoa(i))
	}
	wg.Wait()
}

func TestConcurrentScannerWithWorkerMultiChan(t *testing.T) {
	portCh := make(chan int, 100)
	resultCh := make(chan int)
	defer close(portCh)
	defer close(resultCh)

	for i := 0; i < cap(portCh); i++ {
		// consumer
		go workerMultiChan(i, addr, portCh, resultCh)
	}

	go func() {
		for i := 1; i <= nums; i++ {
			// provider
			portCh <- i
		}
	}()

	var openedPort []int
	for i := 0; i < nums; i++ {
		port := <-resultCh
		if port != 0 {
			openedPort = append(openedPort, port)
		}
	}

	sort.Ints(openedPort)
	for port := range openedPort {
		t.Logf("%d open\n", port)
	}
}
