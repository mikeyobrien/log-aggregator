package collector

import (
	"context"
	"time"
)

type Collector interface {
	Start(ctx context.Context) error
	Close() error
	GetLogs() <-chan LogEntry
}

type LogEntry struct {
	Timestamp time.Time
	Message   string
	Level     string
	Source    string
}
