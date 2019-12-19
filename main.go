package main

import (
	"fmt"
	"os"
)

func main() {
	if err := RunCLI(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
