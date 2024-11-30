package collector

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/fsnotify/fsnotify"
	"github.com/nxadm/tail"
)

type FileCollector struct {
	outChan chan LogEntry
	watcher *fsnotify.Watcher
	path    string
}

func NewFileCollector(path string) (*FileCollector, error) {
	watcher, err := fsnotify.NewWatcher()
	if err != nil {
		log.Fatalln("Failed to create watcher", err)
	}
	return &FileCollector{
		path:    path,
		outChan: make(chan LogEntry, 5),
		watcher: watcher,
	}, nil
}

func (fc *FileCollector) Start(ctx context.Context) error {
	return fc.readLines()
}

func (fc *FileCollector) Close(ctx context.Context) error {
	fc.watcher.Close()
	return nil
}

func (fc *FileCollector) GetLogs() <-chan LogEntry {
	return nil
}

func (fc *FileCollector) readLines() error {
	t, err := tail.TailFile(fc.path, tail.Config{
		Follow: true, // Continue watching
		ReOpen: true, // Reopen the file if it's rotated
		Location: &tail.SeekInfo{ // Start at the end
			Offset: 0,
			Whence: 2,
		},
	})
	if err != nil {
		fmt.Println("Error:", err)
		return err
	}

	// Read new lines as they're written
	for line := range t.Lines {
		if line.Err != nil {
			fmt.Println("Error:", line.Err)
			continue
		}
		fmt.Printf("New line: %s\n", line.Text)
		fc.writeToChan(line.Text)
	}
	return nil
}

func (fc *FileCollector) writeToChan(logline string) error {
	logEntry := LogEntry{
		Timestamp: time.Now(),
		Message:   logline,
		Level:     "",
		Source:    "",
	}
	select {
	case fc.outChan <- logEntry:
		log.Println("DEBUG: Wrote to chan")
	default:
		log.Println("DEBUG: Channel blocked")
	}
	return nil
}
