package main

import (
	"github.com/ngtrdai197/go-rabbitmq/cmd"
	"github.com/ngtrdai197/go-rabbitmq/pkg/logger"
)

func init() {
	logger.InitGlobalLogger()
}

func main() {
	cmd.Execute()
}
