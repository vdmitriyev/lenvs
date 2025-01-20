package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/joho/godotenv"
)

func init() {
}

func main() {
	debugPtr := flag.Bool("debug", false, "print all imported environment variables")
	versionPtr := flag.Bool("version", false, "print info about the current version")
	versionOnlyPtr := flag.Bool("version-only", false, "print the current version only")
	envFilePtr := flag.String("envfile", ".env", "path to .env file")

	flag.Parse()

	if *versionOnlyPtr {
		fmt.Println(version)
		os.Exit(0)
	}

	if *versionPtr {
		fmt.Println("Version: ", version)
		fmt.Println("Build: ", build)
		os.Exit(0)
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
		os.Exit(0) // Exit after debugging
	}
}
