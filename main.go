package main

import (
	"time"

	"github.com/agnaite/1-4-24/gameplay"
)

func main() {

	gameplay := gameplay.New()

	for {
		err := gameplay.Play()

		if err != nil && err.Error() == "finished" {
			break
		}

		time.Sleep(time.Second)
	}
}
