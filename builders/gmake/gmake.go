package gmake

import (
	"os"
	"os/exec"
	"path/filepath"
	"strings"
)

type GMake struct{}

func (gm *GMake) GetName() string {
	return "GNU Make"
}

func (gm *GMake) getBinary() string {
	if f, err := exec.LookPath("gmake"); err == nil {
		return f
	}
	if f, err := exec.LookPath("make"); err == nil {
		return f
	}
	return ""
}

func (gm *GMake) Detect(inputDir string) bool {
	if _, err := os.Stat(filepath.Join(inputDir, "Makefile")); err != nil {
		return false
	}

	bin := gm.getBinary()
	if bin == "" {
		return false
	}

	// binary is GNU Make?
	if out, err := exec.Command(bin, "--version").Output(); err != nil || !strings.Contains(string(out), "GNU") {
		return false
	}

	// there's a github-webhook target?
	return exec.Command(bin, "--dry-run", "--directory", inputDir, "website-builder").Run() == nil
}

func (gm *GMake) Build(inputDir string, outputDir string) *exec.Cmd {
	cmd := exec.Command(gm.getBinary(), "--directory", inputDir, "website-builder")
	cmd.Env = []string{
		"OUTPUT_DIR=" + outputDir,
	}
	return cmd
}
