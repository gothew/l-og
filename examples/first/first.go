package main

import (
	log "github.com/gothew/l-og"
)

func startExample() {
  log.Debug("test")
  log.Info("info message");
  log.Warn("warn message")
}

func main() {
  startExample()
}
