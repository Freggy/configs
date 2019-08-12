// +build mage

package main

import (
	"fmt"
	"github.com/magefile/mage/mg"
	"github.com/magefile/mage/sh"
	"os"
	"path/filepath"
)

func Build() error {
	commit := getCommit()
	tag := getTag()
	name := "BIN" + tag + "-" + commit

	if err := os.Mkdir("target", 0700); err != nil && !os.IsExist(err) {
		return fmt.Errorf("Could not create target directory: %v", err)
	}

	path := filepath.Join("target", name)
	return sh.Run(mg.GoCmd(), "build", "-o", path, "-ldflags="+createFlags(), "main.go")
}

func createFlags() string {
	tag := getTag()
	commit := getCommit()
	return fmt.Sprintf(`-X "main.commit=%s" -X "main.tag=%s"`, commit, tag)
}

func getCommit() string {
	commit, _ := sh.Output("git", "rev-parse", "-q", "--short", "HEAD")
	return commit
}

func getTag() string {
	tag, _ := sh.Output("git", "describe", "--tags")

	if tag == "" {
		tag = "dev"
	}

	return tag
}

