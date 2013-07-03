package main

import (
	"flag"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
)

var (
	inAddr  = flag.String("ingester", "localhost:4263", "UDP server address")
	srvAddr = flag.String("server", "localhost:9090", "HTTP server address")
	workers = flag.Int("workers", 5, "Number of store workers")
)

func init() {
	flag.Parse()
	log.SetOutput(os.Stdout)
}

func main() {
	visits := make(chan Visit)
	store := NewUVStore()

	for i := 0; i < *workers; i++ {
		worker := NewWorker(i, visits, store)
		worker.Start()
	}

	ingester, err := NewIngester(*inAddr, visits)
	if err != nil {
		log.Fatal(err)
	}
	ingester.Start()

	http.HandleFunc("/users", UVHandler(store))

	s := make(chan os.Signal)
	go func() {
		for sig := range s {
			switch sig {
			case syscall.SIGINT, syscall.SIGTERM:
				log.Printf("Caught signal %d: shutdown", sig)
				ingester.Close()
				os.Exit(0)
			default:
				log.Printf("Caught signal %d: ignored", sig)
			}
		}
	}()
	signal.Notify(s)

	err = http.ListenAndServe(*srvAddr, nil)
	if err != nil {
		log.Fatal(err)
	}
}
