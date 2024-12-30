package main

import (
	"bufio"
	"fmt"
	"os"
	"os/exec"
	"strings"
)

// Ensures gofmt doesn't remove the "fmt" import in stage 1 (feel free to remove this!)
var _ = fmt.Fprint

func main() {
	
	reader := bufio.NewReader(os.Stdin)
	for {
		fmt.Fprint(os.Stdout, "$ ")
		cmd, _ := reader.ReadString('\n')
		cmd = strings.TrimSpace(cmd)
		_, err := exec.LookPath(cmd)

		if err != nil {
			fmt.Printf("%s: command not found\n", cmd)
			// os.Exit(1)
		}

	}	
}
