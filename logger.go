package littleLogger

import (
	"bufio"
	"sync"
)

type Logger struct {
	target bufio.Writer
	mu sync.Mutex
}

func NewLogger(target bufio.Writer) *Logger {
	return &Logger{target: target}
}


