package jekyll

import (
	"bufio"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
)

type Jekyll struct {
	bundlerVersion string
}

func (j *Jekyll) GetName() string {
	return "Jekyll"
}

func (j *Jekyll) Detect(inputDir string) bool {
	if _, err := os.Stat(filepath.Join(inputDir, "_config.yml")); err != nil {
		return false
	}

	f, err := os.Open(filepath.Join(inputDir, "Gemfile.lock"))
	if err != nil {
		return false
	}
	defer f.Close()

	found := false
	next := false
	j.bundlerVersion = ""

	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		if next { // this is bundler version
			j.bundlerVersion = strings.TrimSpace(scanner.Text())
			next = false
			continue
		}

		if t := scanner.Text(); strings.Contains(t, "jekyll") || strings.Contains(t, "github-pages") {
			found = true
		}

		if strings.TrimSpace(scanner.Text()) == "BUNDLED WITH" {
			next = true
		}

	}

	return j.bundlerVersion != "" && found
}

func (j *Jekyll) Build(inputDir string, outputDir string) []*exec.Cmd {
	rv := []*exec.Cmd{}

	if strings.HasPrefix(j.bundlerVersion, "1.") {
		cmd := exec.Command("bundle", "install", "--deployment")
		cmd.Dir = inputDir
		rv = append(rv, cmd)
	} else if strings.HasPrefix(j.bundlerVersion, "2.") {
		cmd := exec.Command("bundle", "config", "set", "--local", "deployment", "true")
		cmd.Dir = inputDir
		rv = append(rv, cmd)

		cmd = exec.Command("bundle", "install")
		cmd.Dir = inputDir
		rv = append(rv, cmd)
	} else {
		return nil
	}

	cmd := exec.Command("bundle", "exec", "jekyll", "build", "-V", "-d", outputDir)
	cmd.Dir = inputDir

	return append(rv, cmd)
}
