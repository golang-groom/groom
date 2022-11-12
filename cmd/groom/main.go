package main

import (
	"flag"
	"fmt"
	"os"
	"os/exec"
	"syscall"

	"github.com/pspiagicw/colorlog"
)

func main(){

    version := flag.Bool("version" , false, "Print version info" )
    verbose := flag.Bool("verbose", false, "Print verbose info")

    flag.Parse()

    if *version {
        os.Exit(0)
    }

    if len(flag.Args()) == 0 {
        colorlog.LogError("No subcommand provided")

        fmt.Println("Usage of `groom`")
        fmt.Println("groom [SUBCOMMAND] [OPTIONS]")

        flag.PrintDefaults()
        os.Exit(0)
    }

    arg := flag.Arg(0)

    name := fmt.Sprintf("groom-%s", arg)

    binary , err := exec.LookPath(name)

    if err != nil {
        colorlog.LogFatal("'%s', no such binary found on your $PATH variable", name)
    }

    args := flag.Args()

    env := os.Environ()

    if *verbose {
        env = append(env, "GROOM_VERBOSE=1")
    }

    execErr := syscall.Exec(binary , args, env)

    if execErr != nil {
        colorlog.LogFatal("Error executing subcommand, %q", execErr)
    }

}
