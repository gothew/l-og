package main

import (
	log "github.com/gothew/l-og"
)

func startExample() {
  log.SetLevel(log.DebugLevel)
  log.Debug("test")
  log.Debug("test")
  log.Debug("test")
  log.Debug("test")
  log.Debug("test")
}

func main() {
  startExample()
}
