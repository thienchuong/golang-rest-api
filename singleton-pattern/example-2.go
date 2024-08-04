package main

import (
	"fmt"
	"sync"
)

type Logger struct {
	logs []string
}

var (
	instance *Logger
	once     sync.Once
)

func GetInstance() *Logger {
	once.Do(func() {
		instance = &Logger{}
		instance.logs = make([]string, 0)
	})
	return instance
}

func (l *Logger) Log(message string) {
	l.logs = append(l.logs, message)
}

func (l *Logger) PrintLogs() {
	fmt.Println("Log Message ahiih:")
	for _, log := range l.logs {
		fmt.Println(log)
	}
}

func main() {
	logger := GetInstance()
	logger.Log("Application started")
	logger.Log("Processing data...")
	logger.Log("Application finished")
	logger.PrintLogs()

}
