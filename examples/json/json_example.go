package main

import (
	log "github.com/gothew/l-og"
)

func startExample() {
  log.SetFormatter(log.JSONFormatter)
  log.Info("info message");
}

func main() {
  startExample()
}
