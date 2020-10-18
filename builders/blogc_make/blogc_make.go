package blogc_make

import (
	"os"
	"os/exec"
	"path/filepath"
)

type BlogcMake struct {
	bin string
}

func (bm *BlogcMake) GetName() string {
	return "blogc-make"
}

func (bm *BlogcMake) Available() bool {
	if bm.bin != "" {
		return true
	}
	if b, err := exec.LookPath("blogc-make"); err != nil {
		return false
	} else {
		bm.bin = b
		return true
	}
}

func (bm *BlogcMake) Detect(inputDir string) bool {
	if !bm.Available() {
		return false
	}

	_, err := os.Stat(filepath.Join(inputDir, "blogcfile"))
	return err == nil
}

func (bm *BlogcMake) Build(inputDir string, outputDir string) []*exec.Cmd {
	if !bm.Available() {
		return nil
	}

	cmd := exec.Command(bm.bin, "all")
	cmd.Dir = inputDir
	cmd.Env = []string{
		"OUTPUT_DIR=" + outputDir,
	}
	return []*exec.Cmd{cmd}
}
