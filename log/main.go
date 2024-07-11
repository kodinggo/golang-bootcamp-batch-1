package main

import (
	log "github.com/sirupsen/logrus"
)

func main() {
	log.SetFormatter(&log.JSONFormatter{})

	log.Debug("Ini debug")
	log.Info("Ini Info")
	log.Warn("Ini Warn")
	log.Error("Ini Error")

	log.WithFields(log.Fields{
		"ctx":       "ini context",
		"req":       "request payload",
		"requester": "token/userId/email",
		"ip":        "127.0.0.1",
	}).Info("A user has been deleted")

	log.Fatal("Ini Fatal")
}
