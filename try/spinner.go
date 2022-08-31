package main

import (
	"github.com/Wulfheart/simple-server-automation/progress"
	"time"
)

func main() {
	bar := progress.AddSpinner(&progress.Spinner{
		TimeStarted: time.Now(),
	})
	p := progress.New()

	p.AddSpinner(bar)
	p.Start()

	go p.Listen()

	time.Sleep(4 * time.Second)
	p.Stop()
}
