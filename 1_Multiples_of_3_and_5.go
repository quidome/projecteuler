/*
If we list all the natural numbers below 10 that are multiples of 3 or 5, we
get 3, 5, 6 and 9. The sum of these multiples is 23.

Find the sum of all the multiples of 3 or 5 below 1000.
*/

package main

import (
	"log"
	"math"
	"os"
	"time"
)

func main() {
	// measure run time
	start := time.Now()

	upper := 1000

	// my solution
	sum := 0
	for i := 1; i < upper; i++ {
		if math.Mod(float64(i), 3) == 0 || math.Mod(float64(i), 5) == 0 {
			sum += i
		}
	}
	// present results
	log.Printf("the answer to this quest is: %d", sum)

	// report run time
	log.Printf("%s took %s", os.Args[0], time.Since(start))
}
