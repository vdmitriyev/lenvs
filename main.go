package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"os/exec"
	"runtime"
	"strings"

	"github.com/joho/godotenv"
)

func main() {

	flag.Usage = func() {
		fmt.Println("`lenvs` starts a sub-process with an operating system specific terminal (e.g., cmd, bash, etc.).")
		fmt.Println("The terminal will get environment variables read from an '.env' file.")
		fmt.Println("\nArguments:")
		flag.PrintDefaults()
		fmt.Println("\nExample usage:")
		fmt.Println("\tlenvs .env")

	}

	debugPtr := flag.Bool("debug", false, "print all imported environment variables")
	versionPtr := flag.Bool("version", false, "print info about the current version")
	versionOnlyPtr := flag.Bool("version-only", false, "print the current version only")
	envFilePtr := flag.String("envfile", ".env", "path to .env file")

	flag.Parse()

	if *versionOnlyPtr {
		fmt.Println(version)
		return
	}

	if *versionPtr {
		fmt.Println("Version: ", version)
		fmt.Println("Build: ", build)
		return
	}

	// If no --envfile is provided and there are command-line arguments,
	// assume the first argument is the path to the .env file
	if *envFilePtr == ".env" && flag.NArg() > 0 {
		*envFilePtr = flag.Arg(0)
	}

	if *envFilePtr != ".env" {
		log.Printf("Loading environment from custom file: %s\n", *envFilePtr)
	}

	// Load environment from specified file
	if err := godotenv.Load(*envFilePtr); err != nil {
		log.Printf("No file with environment variables found at: %s\n", *envFilePtr)
	}

	if *debugPtr {
		log.Printf("Debug mode is on. Printing all environment variables\n")
		for _, e := range os.Environ() {
			pair := strings.SplitN(e, "=", 2)
			fmt.Printf("%s = %s\n", pair[0], pair[1])
		}
		return
	}

	var cmd string
	if runtime.GOOS == "windows" {
		cmd = "cmd"
	} else {
		cmd = "bash"
	}

	command := exec.Command(cmd)
	command.Stdin = os.Stdin
	command.Stdout = os.Stdout
	command.Stderr = os.Stderr
	command.Run()
}
