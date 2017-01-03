package main

import (
	"fmt"
	"log"
	"os"
	"time"
)

func l() {
	hostname, _ := os.Hostname()
	host := "[" + hostname + "] "
	multiline := ""
	for i := 0; i < 10; i++ {
		log.Println(host+"log line %d", i)
		multiline += host + fmt.Sprintf("multiline part %d", i) + "\n"
	}
	multiline += host + "end of multiline"
	log.Println(multiline)
}

func main() {
	for {
		l()
		time.Sleep(1300 * time.Millisecond)
	}
}
