
import (
	"log"
	"time"
)

// clear the noise on log output
func fn() {
	log.Println("----------\n")
	defer log.Println("----------\n")

	// ...
}

// create a timer
func StartTimer(name string) func() {
	t := time.Now()
	log.Println(name, "started")
	return func() {
		d := time.Now.Sub(t)
		log.Println(name, "took", d)
	}
}

// example usage of timer
func TimedFunc() {
	stop := StartTimer("TimedFunc")
	defer stop()

	// ...
}

// Accept interfaces, return structs
