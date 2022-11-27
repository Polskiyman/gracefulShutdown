package main

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"
)

func main() {
	intChan := make(chan int)
	signalChan := make(chan os.Signal)
	signal.Notify(signalChan, os.Interrupt, syscall.SIGTERM)
	v := 1
	for {
		v++
		if v == 100 {
			v = 1
		}
		go func() {
			intChan <- v * v
			<-signalChan
			fmt.Println("Выхожу из программы")
			os.Exit(1)
		}()
		fmt.Println(<-intChan)

	}
}
