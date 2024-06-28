package internal
import (
	"fmt"
	"os"
	"os/signal"
	"syscall"
)
func Yes () (err error){
	word:= "y"
	if len(os.Args) > 2 {
		word = os.Args[1]
	}
	// Creating a channel to receive OS signals
	sig := make(chan os.Signal, 1)
	// Notifying the signal channel on receiving Interrupt or Terminate signals
	signal.Notify(sig, syscall.SIGINT, syscall.SIGTERM)
	for {
		if len(sig) != 0 {
			break
		}
		fmt.Println(word)
	}
	return nil
}