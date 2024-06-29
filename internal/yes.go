package internal

import (
	"fmt"
	"os"
	"os/signal"
	"strings"
	"syscall"
	"time"
)

func Yes() (err error) {
	word := "y"
	if len(os.Args) >= 2 {
		word = strings.Join(os.Args[1:], " ")
	}
	// Creating a channel to receive OS signals
	sig := make(chan os.Signal, 1)
	// Notifying the signal channel on receiving Interrupt or Terminate signals
	signal.Notify(sig, syscall.SIGINT, syscall.SIGTERM)
	// Creating a ticker to control the rate of output
	ticker := time.NewTicker(time.Millisecond)
	defer ticker.Stop()
	for {
		select {
		case <-sig:
			return nil
		case <-ticker.C:
			fmt.Println(word)
		}
	}
	// for {
	// 	if len(sig) != 0 {
	// 		break
	// 	}
	// 	fmt.Println(word)
	// 	time.Sleep(1*time.Millisecond)
	// }
	// return nil
}
