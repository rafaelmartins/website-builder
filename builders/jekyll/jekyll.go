package jekyll

import (
	"io/ioutil"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
)

type Jekyll struct{}

func (j *Jekyll) GetName() string {
	return "Jekyll"
}

func (j *Jekyll) Detect(inputDir string) bool {
	if _, err := os.Stat(filepath.Join(inputDir, "_config.yml")); err != nil {
		return false
	}

	// we need Gemfile.lock for bundle deployment mode
	if content, err := ioutil.ReadFile(filepath.Join(inputDir, "Gemfile.lock")); err == nil {
		if strings.Contains(string(content), "jekyll") || strings.Contains(string(content), "github-pages") {
			return true
		}
	}

	return false
}

func (j *Jekyll) Build(inputDir string, outputDir string) []*exec.Cmd {
	cmd1 := exec.Command("bundle", "config", "set", "deployment", "true")
	cmd1.Dir = inputDir

	cmd2 := exec.Command("bundle", "install")
	cmd2.Dir = inputDir

	cmd3 := exec.Command("bundle", "exec", "jekyll", "build", "-d", outputDir)
	cmd3.Dir = inputDir

	return []*exec.Cmd{cmd1, cmd2, cmd3}
}
