package storage

import (
	"time"

	"github.com/mikeyobrien/log-aggregator/internal/collector"
)

type Storage interface {
	Store(entry collector.LogEntry) error
	Search(s string) ([]string, error)
}

// TODO: move this to a search package
type TimeRange struct {
	Start time.Time
	End   time.Time
}

type Query struct {
	TimeRange   TimeRange
	Source      string
	Level       string
	Pattern     string
	ServiceName string
	Limit       int
}
