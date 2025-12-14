package main

import (
	"bufio"
	"fmt"
	"log"
	"os"

	"github.com/ramtinJ95/dbt-lsp/rpc"
)

func main() {
	fmt.Println("hi")
	logger := getLogger("/home/ramtinj/personal-workspace/dbt-lsp/log.txt")
	logger.Printf("hey i started")

	scanner := bufio.NewScanner(os.Stdin)
	scanner.Split(rpc.Split)

	for scanner.Scan() {
		msg := scanner.Text()
		handleMessage(logger, msg)
	}
}

func handleMessage(logger *log.Logger, msg any) {
	logger.Println(msg)
}

func getLogger(filename string) *log.Logger {
	logfile, err := os.OpenFile(filename, os.O_CREATE|os.O_TRUNC|os.O_WRONLY, 0o666)
	if err != nil {
		panic("hey, you didnt give me a good file")
	}

	return log.New(logfile, "[dbt-lsp]", log.Ldate|log.Ltime|log.Lshortfile)
}
