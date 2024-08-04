package main

import (
	"singleton"
)

func main() {
	logger := singleton.GetInstance()
	logger.Log("First message")
	logger.Log("Second message")
	logger.Log("Third message")
	logger.PrintLogs()
}
