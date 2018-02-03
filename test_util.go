//
// test_util.go
// Copyright (C) 2017 Akash Kothawale <akash@decached.com>
//
// Distributed under terms of the MIT license.
//

package git

import (
	"fmt"
	"os"
	"path/filepath"
	"reflect"
	"testing"
)

var test_dir, _ = filepath.Abs("./tmp")

func Setup(testDir string) *Repository {
	path, _ := filepath.Abs(testDir)
	os.Setenv("GIT_DIR", filepath.Join(path, ".git"))
	r, err := Init()
	checkErr(err)
	return r
}

func TearDown(testDir string) {
	path, _ := filepath.Abs(testDir)
	os.RemoveAll(path)
}

// Test Helpers
func assert(got interface{}, want interface{}, t *testing.T) {
	if !reflect.DeepEqual(got, want) {
		t.Error(fmt.Sprintf("Want %s, Got %s", want, got))
	}
}
