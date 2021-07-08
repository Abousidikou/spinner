package spinner

import (
	"fmt"
	"time"
)

func Spinner(delay time.Duration, stop time.Duration) {
	start := time.Now()
	for {
		for _, x := range `-\|/` {
			fmt.Printf("\r%c", x)
			time.Sleep(delay)
		}
		if time.Since(start) == stop {
			break
		}
	}
}
