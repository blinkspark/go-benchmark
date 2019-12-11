package time

import (
	"time"
)

// Runner for a benchmar func
type Runner func()

func Time(r Runner) time.Duration {
	bf := time.Now()
	r()
	af := time.Now()
	return af.Sub(bf)
}
