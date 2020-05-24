package builders

import (
	"errors"
	"os/exec"

	"github.com/rafaelmartins/website-builder/builders/blogc_make"
	"github.com/rafaelmartins/website-builder/builders/gmake"
	"github.com/rafaelmartins/website-builder/builders/jekyll"
)

type Builder interface {
	GetName() string
	Detect(inputDir string) bool
	Build(inputDir string, outputDir string) []*exec.Cmd
}

var builders = []Builder{
	&blogc_make.BlogcMake{},
	&gmake.GMake{},
	&jekyll.Jekyll{},
}

func Detect(inputDir string) (Builder, error) {
	for _, builder := range builders {
		if builder.Detect(inputDir) {
			return builder, nil
		}
	}
	return nil, errors.New("builders: no builder detected")
}
