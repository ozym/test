package main

import (
	"flag"
	"fmt"
	"log"
	"os"
)

func main() {

	flag.Usage = func() {
		fmt.Fprintf(os.Stderr, "\n")
		fmt.Fprintf(os.Stderr, "Provide NTRIP configuration file hiera settings\n")
		fmt.Fprintf(os.Stderr, "\n")
		fmt.Fprintf(os.Stderr, "Usage:\n")
		fmt.Fprintf(os.Stderr, "\n")
		fmt.Fprintf(os.Stderr, "  %s [options]\n", os.Args[0])
		fmt.Fprintf(os.Stderr, "\n")
		fmt.Fprintf(os.Stderr, "General Options:\n")
		fmt.Fprintf(os.Stderr, "\n")
		flag.PrintDefaults()
	}

	var base string
	flag.StringVar(&base, "base", ".", "delta config files")

	var input string
	flag.StringVar(&input, "input", ".", "input csv files")

	var output string
	flag.StringVar(&output, "output", "config.yaml", "output config file")

	flag.Parse()

	config, err := BuildConfig(base, input)
	if err != nil {
		log.Fatalf("unable to build config: %v", err)
	}

	config.Sort()

	if err := config.Write(output); err != nil {
		log.Fatalf("unable to write config file %s: %v", output, err)
	}
}
