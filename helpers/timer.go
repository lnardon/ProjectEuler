package helpers

import (
	"fmt"
	"time"
)

func MeasureTimeSpent() func() {
	start := time.Now()
	return func() {
		fmt.Println("Code executed in:", time.Since(start))
	}
}
