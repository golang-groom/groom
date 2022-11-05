package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"os/exec"
	"syscall"
)

func main(){

    flag.Parse()

    arg := flag.Arg(0)

    if arg == "-h" || arg == "help" {
        flag.PrintDefaults()
    } else {
        binary , err := exec.LookPath(fmt.Sprintf("gurm-%s", arg))

        if err != nil {
            log.Fatalf("No such subcommand found")
        }

        args := flag.Args()[1:]

        env := os.Environ()

        execErr := syscall.Exec(binary , args, env)

        if execErr != nil {
            log.Fatalf("Error executing subcommand")
        }

    }

}
