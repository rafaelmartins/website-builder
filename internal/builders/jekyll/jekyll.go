package jekyll

import (
	"bufio"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
)

type Jekyll struct {
	bin            string
	bundlerVersion string
}

func (j *Jekyll) GetName() string {
	return "Jekyll"
}

func (j *Jekyll) Available() bool {
	if j.bin != "" {
		return true
	}
	if b, err := exec.LookPath("bundle"); err != nil {
		return false
	} else {
		j.bin = b
		return true
	}
}

func (j *Jekyll) Detect(inputDir string) bool {
	if !j.Available() {
		return false
	}

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
	if !j.Available() {
		return nil
	}

	rv := []*exec.Cmd{}

	if strings.HasPrefix(j.bundlerVersion, "1.") {
		cmd := exec.Command(j.bin, "install", "--deployment")
		cmd.Dir = inputDir
		rv = append(rv, cmd)
	} else if strings.HasPrefix(j.bundlerVersion, "2.") {
		cmd := exec.Command(j.bin, "config", "set", "--local", "deployment", "true")
		cmd.Dir = inputDir
		rv = append(rv, cmd)

		cmd = exec.Command(j.bin, "install")
		cmd.Dir = inputDir
		rv = append(rv, cmd)
	} else {
		return nil
	}

	cmd := exec.Command(j.bin, "exec", "jekyll", "build", "-V", "-d", outputDir)
	cmd.Dir = inputDir

	return append(rv, cmd)
}
