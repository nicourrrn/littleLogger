package littleLogger

import (
	"errors"
	"fmt"
	"io"
	"strings"
	"sync"
)

type Logger struct {
	target io.Writer
	mu sync.Mutex
	levels    [4]bool
	formatter func()string
}

//NewLogger lvls = {
//	1 - debug
//	2 - warning TODO Рассмотреть порядок еще раз
//	3 - info
//	4 - error
//}
//	formatter: func()string TODO проверить на деле удобность
//
//
func NewLogger(target io.Writer, lvls ...int) (*Logger, error) {
	if len(lvls) > 4 {
		return nil, errors.New("Count of lvls less than 5")
	}
	levels := [4]bool{false,true,true,true}
	for i, l := range lvls{
		levels[i] = l == 1
	}
	_, err := fmt.Fprint(target, "Starting logs")
	if err != nil {
		return nil, err
	}
	standartFormatter := func() string {return "%msg"}
	return &Logger{target: target, levels: levels, formatter: standartFormatter}, nil
}

func (l *Logger) Error(msg string) {
	fmt.Fprintf(l.target,
		strings.Replace(l.formatter(), "%msg","Error message: %s", 1),
		msg)
}

func (l *Logger) Info(msg string) {
	fmt.Fprintf(l.target,
		strings.Replace(l.formatter(), "%msg","Info message: %s", 1),
		msg)
}

func (l *Logger) Warning(msg string) {
	fmt.Fprintf(l.target,
		strings.Replace(l.formatter(), "%msg","Warning message: %s", 1),
		msg)
}

func (l *Logger) Debug(msg string) {
	fmt.Fprintf(l.target,
		strings.Replace(l.formatter(), "%msg","Debug message: %s", 1),
		msg)
}
