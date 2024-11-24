# Go Event Hub Logger
A lightweight and extensible logging library for Go, designed to log messages to various sinks such as the console or Azure Event Hub. The library supports:
 - Log level filtering (e.g., DEBUG, INFO, WARN, ERROR).
 - JSON-formatted logs.
 - Multiple sinks for versatile logging destinations.

# Features
 - Log Level Filtering: Only log messages at or above the specified log level (e.g., INFO).
 - JSON Logs: Log entries are formatted as JSON for compatibility with modern logging systems.
 - Extensible Design: Add custom sinks easily.
 - Azure Event Hub Integration: Send logs directly to Azure Event Hub for centralized processing.

# Key Components
- logger.Logger:
The main logger that accepts log messages and forwards them to all configured sinks.

- logger.Sink:
An interface for log destinations (e.g., console, Event Hub). You can implement this to create custom sinks.

- logger.LogEntry:
A structured log message that includes a log level, timestamp, message, and optional fields.

- Sinks:
Console Sink (sink_console.go): Outputs logs to the terminal.
Event Hub Sink (sink_eventhub.go): Sends logs to Azure Event Hub.
