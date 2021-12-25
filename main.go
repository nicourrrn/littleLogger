package main

import (
	"fmt"
	littleLogger "github.com/nicourrrn/littleLogger/model/logger"
	"os"
	"time"
)

func main() {
	logger, err := littleLogger.NewLogger(os.Stdout)
	if err != nil {
		fmt.Println("error")
		return
	}
	logger.SetFormatter(func() string {
		return fmt.Sprintf("From %s: $msg", time.Now().String())
	})
	logger.Info("Oh, iam work")
	logger.SetFormatter(littleLogger.FormatterClassic)
	logger.Warning("You use classic")
	logger.Wait()
}
