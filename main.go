package main

import (
	"log"
	"os"
	"path/filepath"

	"github.com/rafaelmartins/website-builder/builders"
	"github.com/rafaelmartins/website-builder/internal/exec"
	"github.com/rafaelmartins/website-builder/internal/symlink"
)

func main() {
	log.SetPrefix("====> ")
	log.SetFlags(0)

	if l := len(os.Args); l < 3 || l > 4 {
		log.Fatalln("error: invalid number of arguments")
	}

	inputDir, err := filepath.Abs(os.Args[1])
	if err != nil {
		log.Fatalln("error:", err)
	}
	outputDir, err := filepath.Abs(os.Args[2])
	if err != nil {
		log.Fatalln("error:", err)
	}

	builder, err := builders.Detect(inputDir)
	if err != nil {
		log.Fatalln("error:", err)
	}

	log.Println("builder:", builder.GetName())

	cmd := builder.Build(inputDir, outputDir)
	if cmd == nil {
		log.Fatalln("error: command not found")
	}

	status, err := exec.Run(cmd)
	if err != nil {
		if status == 0 {
			log.Fatalln("error:", err)
		}
		os.Exit(status)
	}

	if len(os.Args) == 4 {
		if err := symlink.Update(os.Args[3], os.Args[2]); err != nil {
			log.Fatalln("error:", err)
		}
	}
}
