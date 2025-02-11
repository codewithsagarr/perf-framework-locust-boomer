package boomer

import (
	"fmt"
	"os"
	"os/signal"
	"sync"
	"syscall"

	"github.com/myzhan/boomer"
)

// Used to subscribe to the boomer:quit signal and clean things up once this happens.
func waitForQuit() {
	wg := sync.WaitGroup{}
	wg.Add(1)

	quitByMe := false
	go func() {
		c := make(chan os.Signal, 1)
		signal.Notify(c, syscall.SIGINT, syscall.SIGTERM)
		<-c
		quitByMe = true
		wg.Done()
	}()

	fmt.Println("Subscribe boomer:quit")
	boomer.Events.Subscribe("boomer:quit", func() {
		if !quitByMe {
			wg.Done()
			fmt.Println("boomer:quit")
		}
	})

	wg.Wait()
}
