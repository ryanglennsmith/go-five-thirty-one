package util

import (
	"fmt"
	"time"
)

func StartLoadingIndicator() chan bool {
	done := make(chan bool)
	go func() {
		spinner := []string{"|", "/", "-", "\\"}
		i := 0
		for {
			select {
			case <-done:
				fmt.Print("\r \r")
				return
			default:
				fmt.Printf("\rworking %s", spinner[i%len(spinner)])
				i++
				time.Sleep(100 * time.Millisecond)
			}
		}
	}()
				
	return done
}