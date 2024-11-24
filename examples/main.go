package main

import (
	"go-eventhub-logger/logger"
)

func main() {
	// Create sinks
	consoleSink := &logger.ConsoleSink{}
	//add the connection string to your EventHub
	eventHubSink, err := logger.NewEventHubSink("EventHub Connection String")
	if err != nil {
		panic(err)
	}
	defer eventHubSink.Close()

	// Create logger
	log := logger.NewLogger(logger.Info, consoleSink, eventHubSink)
	defer log.Close()

	// Log messages
	log.Log(logger.Info, "Application started", map[string]interface{}{
		"version": "1.0.0",
	})
	log.Log(logger.Warn, "Low disk space", map[string]interface{}{
		"disk": "C:",
		"free": "2GB",
	})
}
