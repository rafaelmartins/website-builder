package builders

import (
	"errors"
	"os/exec"
	"strings"

	"github.com/rafaelmartins/website-builder/internal/builders/blogc_make"
	"github.com/rafaelmartins/website-builder/internal/builders/gmake"
	"github.com/rafaelmartins/website-builder/internal/builders/jekyll"
)

type Builder interface {
	GetName() string
	Available() bool
	Detect(inputDir string) bool
	Build(inputDir string, outputDir string) []*exec.Cmd
}

type Builders []Builder

var builders = Builders{
	&blogc_make.BlogcMake{},
	&gmake.GMake{},
	&jekyll.Jekyll{},
}

func Available() (Builders, error) {
	rv := Builders{}
	for _, builder := range builders {
		if builder.Available() {
			rv = append(rv, builder)
		}
	}
	if len(rv) == 0 {
		return nil, errors.New("builders: no builder available")
	}
	return rv, nil
}

func (b Builders) Detect(inputDir string) (Builder, error) {
	for _, builder := range b {
		if builder.Detect(inputDir) {
			return builder, nil
		}
	}
	return nil, errors.New("builders: no builder detected")
}

func (b Builders) String() string {
	n := []string{}
	for _, builder := range b {
		n = append(n, builder.GetName())
	}
	return strings.Join(n, ", ")
}
