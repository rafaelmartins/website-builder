package gmake

import (
	"os"
	"os/exec"
	"path/filepath"
	"strings"
)

type GMake struct {
	bin string
}

func (gm *GMake) GetName() string {
	return "GNU Make"
}

func (gm *GMake) getBinary() string {
	if gm.bin != "" {
		return gm.bin
	}
	if f, err := exec.LookPath("gmake"); err == nil {
		gm.bin = f
		return f
	}
	if f, err := exec.LookPath("make"); err == nil {
		gm.bin = f
		return f
	}
	return ""
}

func (gm *GMake) Available() bool {
	bin := gm.getBinary()
	if bin == "" {
		return false
	}

	// binary is GNU Make?
	if out, err := exec.Command(bin, "--version").Output(); err != nil || !strings.Contains(string(out), "GNU") {
		return false
	}

	return true
}

func (gm *GMake) Detect(inputDir string) bool {
	if !gm.Available() {
		return false
	}

	if _, err := os.Stat(filepath.Join(inputDir, "Makefile")); err != nil {
		return false
	}

	// there's a website-builder target?
	return exec.Command(gm.getBinary(), "--dry-run", "--directory", inputDir, "website-builder").Run() == nil
}

func (gm *GMake) Build(inputDir string, outputDir string) []*exec.Cmd {
	if !gm.Available() {
		return nil
	}

	cmd := exec.Command(gm.getBinary(), "--directory", inputDir, "website-builder")
	cmd.Env = []string{
		"OUTPUT_DIR=" + outputDir,
	}
	return []*exec.Cmd{cmd}
}
