package main

import (
	"flag"
	"log"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func main() {
	for {
		name := flag.String("name", "hello world", "name description")
		flag.Parse()
		time.Sleep(1 * time.Second)
		log.Println("Hello world 0.3.5!!!")
		sigs := make(chan os.Signal)
		signal.Notify(sigs, syscall.SIGQUIT)
		go func() {
			s := <-sigs
			log.Printf("RECEIVED SIGNAL: %s", s)
			os.Exit(1)
		}()
		for {
			log.Printf("hello %s", *name)
			time.Sleep(1 * time.Second)
		}
	}
}
