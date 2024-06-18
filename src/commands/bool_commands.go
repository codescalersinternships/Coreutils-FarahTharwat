package commands

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"
)
func Yes (word string){
	// Creating a channel to receive OS signals
	sig := make(chan os.Signal, 1)
	// Notifying the signal channel on receiving Interrupt or Terminate signals
	signal.Notify(sig, syscall.SIGINT, syscall.SIGTERM)
	for {
		if len(sig) != 0 {
			break
		}
		if word == " "{
			fmt.Println("y")
		}else {
			fmt.Println(word)
		}	
	}
}
func True () (int){
	return 0 
}
func False ()(int){
	return 1
}