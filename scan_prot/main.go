package main

import (
	"fmt"
	"net"
	"os"
	"runtime"
	"strconv"
	"sync"
)

func main() {
	runtime.GOMAXPROCS(runtime.NumCPU())
	if len(os.Args) != 2 {
		fmt.Fprintf(os.Stderr, "Usage: %s host ", os.Args[0])
		os.Exit(1)
	}
	addr := os.Args[1]
	fmt.Println("[+]Scaning", addr)
	wg := sync.WaitGroup{}
	gortNums := make(chan struct{}, 100)
	//pro := []string{"80", "3306", "22", "3389", "1433", "8080", "25", "443", "8000", "21"}
	for i := 0; i < 65536; i++ {
		//port := pro[i]
		wg.Add(1)
		gortNums <- struct{}{}
		go ScanPort(addr, strconv.Itoa(i), &wg,gortNums)
	}
	wg.Wait()

}

func ScanPort(ip, port string, wg *sync.WaitGroup,s <-chan struct{}) {
	service := ip + ":" + port
	_, err := net.Dial("tcp", service)
	if err != nil {
		//fmt.Println("[-]", port, "Close")
	} else {
		fmt.Println("[+]", port, "Open")
	}
	<-s
	wg.Done()
}
