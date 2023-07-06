package main

import (
	"time"

	"github.com/go-co-op/gocron"
)

func main() {
	scheduler := gocron.NewScheduler(time.Local)
	if _, err := scheduler.Every(1).Second().Do(func() {
		println("Hello, GO cron")
	}); err != nil {
		panic(err)
	}
	scheduler.StartBlocking()
}
