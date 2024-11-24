package logger

import (
	"context"
	"encoding/json"
	"time"

	eventhub "github.com/Azure/azure-event-hubs-go/v3"
)

type EventHubSink struct {
	client *eventhub.Hub
}

func NewEventHubSink(connectionString string) (*EventHubSink, error) {
	hub, err := eventhub.NewHubFromConnectionString(connectionString)
	if err != nil {
		return nil, err
	}
	return &EventHubSink{client: hub}, nil
}

func (e *EventHubSink) Write(entry LogEntry) error {
	data, err := json.Marshal(entry)
	if err != nil {
		return err
	}
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	return e.client.Send(ctx, eventhub.NewEvent(data))
}

func (e *EventHubSink) Close() error {
	return e.client.Close(context.Background())
}
