package main

import (
	"fmt"
	"log"
	"os"
	"os/signal"
	"sync"

	"github.com/biter777/viz-go-lib"
	"github.com/biter777/viz-go-lib/api"
)

func main() {
	cls, _ := viz.NewClient("wss://solox.world/ws")
	defer cls.Close()

	cls.API.SetBlockAppliedCallback(func(block *api.CallbackBlockResponse, err error) {
		log.Printf("%+v \n", block)
	})

	WaitForCtrlC()
	fmt.Println("END")
}

func WaitForCtrlC() {
	var end_waiter sync.WaitGroup
	end_waiter.Add(1)
	var signal_channel chan os.Signal
	signal_channel = make(chan os.Signal, 1)
	signal.Notify(signal_channel, os.Interrupt)
	go func() {
		<-signal_channel
		end_waiter.Done()
	}()
	end_waiter.Wait()
}
