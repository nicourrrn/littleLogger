package main

import (
	"fmt"
	littleLogger "github.com/nicourrrn/littleLogger/model/logger"
	"os"
)

func main() {
	logger, err := littleLogger.NewLogger(os.Stdout)
	if err != nil {
		fmt.Println("error")
		return
	}
	//logger.SetFormatter(func() string {
	//	return fmt.Sprintf("From: $msg")
	//})
	logger.SetFormatter(littleLogger.FormatterClassic)
	logger.Info("Oh, iam work")
	logger.SetFormatter(littleLogger.FormatterMinimal)
	logger.Warning("You use classic")
	logger.Wait()
}
