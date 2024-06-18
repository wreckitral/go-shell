package main

import (
	"bufio"
	"fmt"
	"os"
	"slices"
	"strings"
)

func main() {
    builtin := []string{"echo", "type", "exit"}

    for {
        fmt.Fprint(os.Stdout, "$ ")

        command, _ := bufio.NewReader(os.Stdin).ReadString('\n')
        if command == "exit 0\n" {
            os.Exit(0)
        }

        switch {
        case strings.HasPrefix(command, "echo "):
            fmt.Printf("%s", strings.TrimLeft(command, "echo "))
        case strings.HasPrefix(command, "type "):
            tool := strings.TrimPrefix(command, "type ")
            tool = strings.TrimSpace(tool)

            if slices.Contains(builtin, tool) {
                fmt.Printf("%s is a shell builtin\n", tool)
            } else {
                fmt.Printf("%s: not found\n", tool)
            }
        default:
            fmt.Printf("%s: command not found\n", strings.TrimRight(command, "\n"))
        }
    }
}

