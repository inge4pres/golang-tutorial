package main

import (
	"fmt"
	"strconv"
	"sync"
)

var wg sync.WaitGroup
var strCh = make(chan (string))

func intToStr(in int) string {
	return fmt.Sprintf("%s", strconv.Itoa(in))
}

func printStr(in string) {
	fmt.Printf("STR: %s\n", in)
}

func main() {

	for i := 0; i < 100; i++ {
		go func(i int) {
			strCh <- intToStr(i)
			wg.Add(1)
		}(i)

		printStr(<-strCh)
		wg.Done()

	}

	wg.Wait()
	return
}
