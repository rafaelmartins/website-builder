package symlink

import (
	"log"
	"os"
	"path/filepath"
)

func Update(symlink string, target string) error {
	absSymlink, err := filepath.Abs(symlink)
	if err != nil {
		return err
	}

	absTarget, err := filepath.Abs(target)
	if err != nil {
		return err
	}

	if oldTarget, err := filepath.EvalSymlinks(absSymlink); err == nil {
		if absTarget == oldTarget {
			return nil
		}

		os.RemoveAll(oldTarget)
		os.Remove(absSymlink)
	}

	relTarget, err := filepath.Rel(filepath.Dir(absSymlink), absTarget)

	log.Printf("====> creating symlink: %s -> %s", absSymlink, relTarget)
	if err := os.MkdirAll(filepath.Dir(absSymlink), 0777); err != nil {
		return err
	}
	if err := os.Symlink(relTarget, absSymlink); err != nil {
		return err
	}

	return nil
}
