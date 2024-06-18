package main

import (
	"bufio"
	"fmt"
	"os"
	"os/exec"
	"strings"
)

var builtin = map[string]int{"echo": 0, "type": 1, "pwd": 2, "exit": 3}

func main() {

    for {
        fmt.Fprint(os.Stdout, "$ ")

        c, _ := bufio.NewReader(os.Stdin).ReadString('\n')
        if c == "exit 0\n" {
            os.Exit(0)
        }

        input := strings.TrimRight(c, "\n")
        splittedInput := strings.Split(input, " ")
        cmd := splittedInput[0]
        args := splittedInput[1:]

        if value, ok := builtin[cmd]; !ok {
            cliApp(cmd, args)
        } else {
            switch value {
            case 0:
                echoCmd(args) 
            case 1:
                typeCmd(args)
            case 2:
                pwdCmd()
            }


        }
    }
}

func echoCmd(args []string) {
    fmt.Printf("%s\n", strings.Join(args, " "))
}

func typeCmd(args []string) {
    if _, ok := builtin[args[0]]; ok {
        fmt.Printf("%s is a shell builtin\n", args[0])
        return
    } else {
        paths := strings.Split(os.Getenv("PATH"), ":")

        for _, path := range paths {
            if _, err := os.Stat(path + "/" + args[0]); err == nil {
                fmt.Printf("%s is %s\n", args[0], path + "/" + args[0])
                return
            }
        }
    }

    fmt.Printf("%s: not found\n", args[0])
}

func pwdCmd() {
    pwd, err := os.Getwd()
    if err != nil {
        fmt.Println(err.Error())
        return
    }

    fmt.Printf("%s\n", pwd)
}

func cliApp(cmd string, args []string) {
    command := exec.Command(cmd, args...)
    command.Stderr = os.Stderr
    command.Stdout = os.Stdout

    if err := command.Run(); err != nil {
        fmt.Printf("%s: command not found\n", cmd)
        return
    }

}
