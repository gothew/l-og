package main

import (
	"os"
	"time"

	log "github.com/gothew/l-og"
)

func main() {
	logger := log.NewWithOptions(os.Stderr, log.Options{
		ReportCaller:    true,
		ReportTimestamp: true,
		TimeFormat:      time.Kitchen,
	})
	logger.Info("Starting oven!", "degree", 375)
	time.Sleep(10 * time.Second)
	logger.Info("Finished baking")
}
