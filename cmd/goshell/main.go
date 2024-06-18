package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
    for {
        fmt.Fprint(os.Stdout, "$ ")
        command, _ := bufio.NewReader(os.Stdin).ReadString('\n')
        fmt.Printf("%s: command not found\n", strings.TrimRight(command, "\n"))
    }
}

