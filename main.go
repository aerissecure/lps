// Count lines of input per second on stdin
package main

import (
	"bufio"
	"fmt"
	"os"
	"time"
	"flag"
)

func readLines(c chan int) {
	count := 0
	bio := bufio.NewReader(os.Stdin)
	for {
		_, more, err := bio.ReadLine()
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
		if more {
			continue
		}
		count += 1
		c <- count
	}
}

func main() {
	var d time.Duration
	flag.DurationVar(&d, "i", time.Second, "Update interval")
	flag.Parse()
	line := 0
	count := 0
	c := make(chan int)
	tick := time.Tick(d)
	go readLines(c)

	for {
		select {
		// print counts
		case <-tick:
			fmt.Println(float64(line-count)/d.Seconds(), "/sec")
			count = line
		// update counts
		case line = <-c:
		}
	}
}
