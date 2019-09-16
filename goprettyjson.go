package main

import (
	"bufio"
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"os"

	"github.com/Jeffail/gabs"
)

func main() {
	info, err := os.Stdin.Stat()
	if err != nil {
		panic(err)
	}

	if info.Mode()&os.ModeCharDevice != 0 || info.Size() <= 0 {
		fmt.Println("The command is intended to work with pipes.")
		fmt.Println("Usage: fortune | goprettyjson")
		return
	}

	reader := bufio.NewReader(os.Stdin)
	var output []byte

	for {
		input, _, err := reader.ReadRune()
		if err != nil && err == io.EOF {
			break
		}
		output = append(output, byte(input))
	}

	jsonParsed, err := gabs.ParseJSON(output)
	if err != nil {
		log.Fatalf("Failed to parse JSON: %s", err)
	}

	var out bytes.Buffer
	json.Indent(&out, jsonParsed.Bytes(), "", "\t")
	if err != nil {
		log.Fatalf("Failed to indent JSON: %s", err)
	}
	out.WriteTo(os.Stdout)
	if err != nil {
		log.Fatalf("Failed to write to standard out: %s", err)
	}
}
