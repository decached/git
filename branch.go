//
// branch.go
// Copyright (C) 2017 Akash Kothawale <akash@decached.com>
//
// Distributed under terms of the MIT license.
//

package git

import (
	// "fmt"
	// 	"io/ioutil"
	"errors"
	"os"
)

func (r *Repository) setHEAD(branch string) error {
	if _, err := os.Stat(r.dir + "/refs/heads/" + branch); branch != "master" && os.IsNotExist(err) {
		return errors.New(BRANCH_NOT_EXIST)
	}
	f, err := os.Create(r.dir + "/HEAD")
	if err != nil {
		return err
	}

	f.WriteString("ref: refs/heads/" + branch + "\n")
	return nil
}

func (r *Repository) Checkout(branch string, options map[string]string) error {
	return r.setHEAD(branch)
}
