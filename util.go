//
// util.go
// Copyright (C) 2017 Akash Kothawale <akash@decached.com>
//
// Distributed under terms of the MIT license.
//

package git

import (
	"crypto/sha1"
	"encoding/hex"
	"os"
)

// Utilities
func hash(r []byte) string {
	sha := sha1.Sum(r)
	return string(hex.EncodeToString(sha[:]))
}

// Helpers
func getenv(k string, default_ string) string {
	v := os.Getenv(k)
	if v == "" {
		return default_
	}
	return v
}

func checkErr(e error) {
	if e != nil {
		panic(e)
	}
}
