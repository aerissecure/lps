// Count lines of input per second on stdin
package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
	"time"
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
	var t bool
	flag.DurationVar(&d, "i", time.Second, "Update interval")
	flag.BoolVar(&t, "t", false, "Include timestamp")
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
			prnt := fmt.Sprintf("%v /sec", float64(line-count)/d.Seconds())
			if t {
				prnt = fmt.Sprintf("%s\t%s", prnt, time.Now().UTC().Format("Mon Jan 2 15:04:05 UTC 2006"))
			}
			fmt.Println(prnt)
			count = line
		// update counts
		case line = <-c:
		}
	}
}
