package main

import (
	"flag"
	"fmt"
	"os"
	"os/exec"
	"runtime"
	"strings"
	"syscall"

	"github.com/pspiagicw/colorlog"
)

// Main function
func main() {

	// Flags for version and verbose
	version := flag.Bool("version", false, "Print version info")
	verbose := flag.Bool("verbose", false, "Print verbosely")
	list := flag.Bool("list", false, "List all subcommands.")

	// parse the flags themselves.
	flag.Parse()

	if *version {
		printVersion()
	}

	if *list {
		listSubcommands()
	}

	if len(flag.Args()) == 0 {
		noSubcommand()
	}

	arg := flag.Arg(0)

	env := os.Environ()

	if *verbose {
		env = append(env, "GROOM_VERBOSE=1")
	}

	executeBinary(env, arg)

}
func executeBinary(env []string, arg string) {

	name := fmt.Sprintf("groom-%s", arg)

	binary, err := exec.LookPath(name)

	if err != nil {
		colorlog.LogFatal("'%s', no such binary found on your $PATH variable", name)
	}

	args := flag.Args()

	execErr := syscall.Exec(binary, args, env)

	if execErr != nil {
		colorlog.LogFatal("Error executing subcommand, %q", execErr)
	}

}
func noSubcommand() {
	colorlog.LogError("No subcommand provided")

	fmt.Println("Usage of `groom`")
	fmt.Println("groom [SUBCOMMAND] [OPTIONS]")

	flag.PrintDefaults()
	os.Exit(0)

}
func printVersion() {
	colorlog.LogInfo("groom version: %s", "1.0")
	os.Exit(0)
}
func listSubcommands() {
	PATH, exists := os.LookupEnv("PATH")

	commands := getSubcommandList(PATH)

	if !exists {
		colorlog.LogFatal("Error accessing PATH variable.")
	}

	if len(commands) == 0 {
		colorlog.LogInfo("There are no groom subcommands installed.")
	} else {
		colorlog.LogSuccess("These are the subcommands available to `groom")
		for command := range commands {
			colorlog.LogInfoFancy("%s", command)

		}

	}

	os.Exit(0)

}
func getSubcommandList(PATH string) map[string]bool {

	pathSplitter := getPathSplitter()

	dirs := strings.Split(PATH, pathSplitter)

	commands := map[string]bool{}

	for _, directory := range dirs {
		entries, err := os.ReadDir(directory)

		if err != nil {
			colorlog.LogError("Error accesing directories on PATH. %q", err)
		}

		for _, entry := range entries {

			// Not a directory
			if !entry.IsDir() {
				name := entry.Name()
				if strings.HasPrefix(name, "groom-") {
					commands[name] = true
				}

			}

		}

	}
	return commands
}
func getPathSplitter() string {
	if runtime.GOOS == "windows" {
		return ";"
	}
	return ":"
}
