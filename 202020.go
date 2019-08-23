package main

import (
	"time"

	al "github.com/gen2brain/beeep"
)

func main() {

	for {
		al.Alert("202020 Rule", "Look at an object at least 20 feet away for at least 20 seconds.", "assets/warning.png")
		time.Sleep(time.Minute * 20)

	}

}
