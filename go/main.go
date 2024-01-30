package main

import (
	"os"

	"github.com/abasnfarah/interpreter/go/repl"
	"go.uber.org/zap"
)

func main() {
	logger, _ := zap.NewProduction()
	logger.Info("Monkey Interpreter in Go")

	repl.Start(os.Stdin, os.Stdout)
}
