package storage

import (
	"sync"
	"time"

	"github.com/mikeyobrien/log-aggregator/internal/collector"
)

type TimeIndex struct {
	entries map[time.Time][]string
	mu      sync.RWMutex
}

type FileStorage struct {
	timeIndex *TimeIndex
	baseDir   string
	indexMu   sync.RWMutex
}

func NewFileStorage(baseDir string, maxLevels uint8) {
}

func (fs *FileStorage) Store(entry collector.LogEntry) error {
	return nil
}
