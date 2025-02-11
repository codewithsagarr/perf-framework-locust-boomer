package boomer

import (
	"time"

	"boomer/globals"
)

// This is kicked off as a go routine to quit boomer when the test duration has completed
func stopStandalone(duration int) {
	time.Sleep(time.Duration(duration) * time.Second)
	globals.GlobalBoomer.Quit()
}
