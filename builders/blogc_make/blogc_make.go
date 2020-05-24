package blogc_make

import (
	"os"
	"os/exec"
	"path/filepath"
)

type BlogcMake struct {
	blogcfile string
}

func (bm *BlogcMake) GetName() string {
	return "blogc-make"
}

func (bm *BlogcMake) Detect(inputDir string) bool {
	_, err := os.Stat(filepath.Join(inputDir, "blogcfile"))
	return err == nil
}

func (bm *BlogcMake) Build(inputDir string, outputDir string) []*exec.Cmd {
	cmd := exec.Command("blogc-make", "all")
	cmd.Dir = inputDir
	cmd.Env = []string{
		"OUTPUT_DIR=" + outputDir,
	}
	return []*exec.Cmd{cmd}
}
