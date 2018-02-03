//
// repository_test.go
// Copyright (C) 2017 Akash Kothawale <akash@decached.com>
//
// Distributed under terms of the MIT license.
//

package git

import (
	"fmt"
	"testing"
)

func TestGetRepository(t *testing.T) {
	Setup("TestGetRepository")
	defer TearDown("TestGetRepository")

	_, _ = GetRepository()
}

func TestInit(t *testing.T) {
	Setup("TestInit")
	defer TearDown("TestInit")

	got, _ := Init()
	fmt.Println(got)
}

func TestAdd(t *testing.T) {
	//r := Setup("TestAdd")
	//defer TearDown("TestAdd")
	//
	//files := []string{"A", "B"}
	//for _, f := range files {
	//	os.Create(filepath.Join(test_dir, f))
	//}
	//_, err := r.Add(files)
	//checkErr(err)
}
