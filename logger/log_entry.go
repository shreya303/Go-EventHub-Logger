package logger

type LogEntry struct {
	Level   string                 `json:"level"`
	Time    string                 `json:"time"`
	Message string                 `json:"message"`
	Fields  map[string]interface{} `json:"fields,omitempty"`
}

type Sink interface {
	Write(entry LogEntry) error
	Close() error
}
