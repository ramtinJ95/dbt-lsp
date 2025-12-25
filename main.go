package main

import (
	"bufio"
	"encoding/json"
	"log"
	"os"

	"github.com/ramtinJ95/dbt-lsp/lsp"
	"github.com/ramtinJ95/dbt-lsp/rpc"
)

func main() {
	logger := getLogger("/home/ramtinj/personal-workspace/dbt-lsp/log.txt")
	logger.Printf("hey i started")

	scanner := bufio.NewScanner(os.Stdin)
	scanner.Split(rpc.Split)

	for scanner.Scan() {
		msg := scanner.Bytes()
		method, contents, err := rpc.DecodeMessage(msg)
		if err != nil {
			logger.Printf("got an err: %s", err)
		}
		handleMessage(logger, method, contents)
	}
}

func handleMessage(logger *log.Logger, method string, contents []byte) {
	logger.Printf("Recived msg with method: %s", method)

	switch method {
	case "initialize":
		var request lsp.InitializeRequest
		if err := json.Unmarshal(contents, &request); err != nil {
			logger.Printf("Hey this was not parsable: %s", err)
		}
		logger.Printf("Connected to: %s %s",
			request.Params.ClientInfo.Name,
			request.Params.ClientInfo.Version)

		msg := lsp.NewInitializeResponse(request.ID)
		reply := rpc.EncodeMessage(msg)
		writer := os.Stdout
		writer.Write([]byte(reply))

		logger.Printf("sent the reply: %s", reply)

	case "textDocument/didOpen":
		var request lsp.DidOpenTextDocumentNotification
		if err := json.Unmarshal(contents, &request); err != nil {
			logger.Printf("hey could not parse this: %s", err)
		}
		logger.Printf("Opened: %s %s ",
			request.Params.TextDocument.URI,
			request.Params.TextDocument.Text)

	}
}

func getLogger(filename string) *log.Logger {
	logfile, err := os.OpenFile(filename, os.O_CREATE|os.O_TRUNC|os.O_WRONLY, 0o666)
	if err != nil {
		panic("hey, you didnt give me a good file")
	}

	return log.New(logfile, "[dbt-lsp]", log.Ldate|log.Ltime|log.Lshortfile)
}
