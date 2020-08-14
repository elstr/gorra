package main

import (
	"flag"
	"fmt"
	"time"
)

func main() {
	mins := flag.Int("m", 2, "minutes to gorrear")
	flag.Parse()
	fmt.Printf("\n\n 👮 Gorreando %d mins 👮 \n", *mins)

	quit := make(chan bool)

	go func() {
		time.Sleep(time.Duration(*mins) * time.Minute)
		quit <- true
	}()

	for {
		select {
		case <- quit:
			fmt.Print("\n 🚨🚨🚨🚨🚨🚨🚨🚨🚨🚨🚨🚨🚨🚨🚨🚨🚨🚨🚨 TERMINO 🚨🚨🚨🚨🚨🚨🚨🚨🚨🚨🚨🚨🚨🚨🚨🚨🚨🚨🚨 ")
			return
		default:
			time.Sleep(700* time.Millisecond)
			fmt.Print(".")
		}
	}
}
