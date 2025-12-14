package main

import (
	"bufio"
	"fmt"
	"os"

	"github.com/ramtinJ95/dbt-lsp/rpc"
)

func main() {
	fmt.Println("hi")

	scanner := bufio.NewScanner(os.Stdin)
	scanner.Split(rpc.Split)

	for scanner.Scan() {
		msg := scanner.Text()
		handleMessage(msg)
	}
}

func handleMessage(_ any) {
}
