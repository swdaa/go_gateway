package main

import (
	"fmt"
	"net"
	"sort"
)

//method 1
// func main() {
// 	for i := 21; i < 120; i++ {
// 		address := fmt.Sprintf("127.31.136.199:%d", i)
// 		conn, err := net.Dial("tcp", address)
// 		if err != nil {
// 			fmt.Printf("%s close \n", address)
// 			continue
// 		}
// 		conn.Close()
// 		fmt.Printf("%s opened \n", address)
// 	}
// }

// //method 2
// func main() {
// 	start := time.Now()
// 	var wg sync.WaitGroup

// 	for i := 21; i < 120; i++ {
// 		wg.Add(1)
// 		go func(port int) {
// 			wg.Done()
// 			address := fmt.Sprintf("127.31.136.199:%d", port)
// 			conn, err := net.Dial("tcp", address)
// 			if err != nil {
// 				fmt.Printf("%s close \n", address)
// 				return
// 			}
// 			conn.Close()
// 			fmt.Printf("%s opened \n", address)
// 		}(i)
// 	}
// 	wg.Wait()
// 	elapsed := time.Since(start) / 1e9
// 	fmt.Printf("\n\n%d seconds", elapsed)
// }

// method 2

func worker(ports chan int, results chan int) {
	for p := range ports {
		address := fmt.Sprintf("20.194.168.28:%d", p)
		conn, err := net.Dial("tcp", address)
		if err != nil {
			fmt.Println("%d error !!!!", p)
			results <- p
			continue
		}
		conn.Close()
		fmt.Printf("%d opened", p)
		wg.Done()
	}
}

func main() {
	ports := make(chan int, 100)
	results := make(chan int)
	var openedports []int
	var closedports []int

	for i := 0; i < cap(ports); i++ {
		go worker(ports, results)
	}

	go func() {
		for i := 1; i < 1024; i++ {
			ports <- i
		}
	}()

	for i := 1; i < 1024; i++ {
		port := <-results
		if port != 0 {
			openedports = append(openedports, port)
		} else {
			closedports = append(closedports, port)
		}
	}

	close(ports)
	close(results)

	sort.Ints(openedports)
	sort.Ints(closedports)
	for _, port := range closedports {
		fmt.Printf("%d", port)
	}
}
