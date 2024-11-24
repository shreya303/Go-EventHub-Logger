package logger

import (
	"encoding/json"
	"fmt"
)

type ConsoleSink struct{}

func (c *ConsoleSink) Write(entry LogEntry) error {
	data, err := json.Marshal(entry)
	if err != nil {
		return err
	}
	fmt.Println(string(data))
	return nil
}

func (c *ConsoleSink) Close() error {
	return nil
}
