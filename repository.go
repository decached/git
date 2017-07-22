//
// repository.go
// Copyright (C) 2017 Akash Kothawale <akash@decached.com>
//
// Distributed under terms of the MIT license.
//

package git

import (
	"os"
	"path/filepath"
)

type Repository struct {
	dir string
}

func GetRepository() (*Repository, error) {
	default_dir, _ := filepath.Abs("./.git")
	base_dir := getenv("GIT_DIR", default_dir)
	return &Repository{
		dir: base_dir,
	}, nil
}

func Init() (*Repository, error) {
	base_dir := getenv("GIT_DIR", ".git")

	dir, err := filepath.Abs(base_dir)
	if err != nil {
		return nil, err
	}

	dirs := []string{
		// "/info",
		// "/logs",
		// "/hooks",
		"/objects",
		// "/objects/info",
		// "/objects/pack",
		"/refs",
		"/refs/heads",
		// "/refs/tags",
	}

	for _, d := range dirs {
		err = os.MkdirAll(base_dir+d, os.ModePerm)
		if err != nil {
			return nil, err
		}
	}

	files := []string{
		"/HEAD",
		// "/config",
		// "/description"
	}
	for _, f := range files {
		_, err = os.Create(base_dir + f)
		if err != nil {
			return nil, err
		}
	}

	r := &Repository{dir: dir}
	r.Checkout("master")

	return r, nil
}
