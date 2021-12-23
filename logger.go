package littleLogger

import (
	"io"
	"sync"
	"fmt"
	"errors"
)


type Logger struct {
	target io.Writer
	mu sync.Mutex
	levels [4]bool
}

//NewLogger lvls = {
//	1 - debug
//	2 - warning TODO Рассмотреть порядок еще раз
//	3 - info
//	4 - error
//}
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
	return &Logger{target: target, levels: levels}, nil
}


