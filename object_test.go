//
// blob_test.go
// Copyright (C) 2017 Akash Kothawale <akash@decached.com>
//
// Distributed under terms of the MIT license.
//

package git

import (
	"fmt"
	"testing"
)

var want_body = []byte("license:mit")
var want_sha = "ab14967724c94d1d5c29ba16bd4505997f07b392"
var want_type = "blob"

func TestGetObject(t *testing.T) {
	r := Setup("TestGetObject")
	defer TearDown("TestGetObject")

	r.WriteObject(want_body, want_type)

	want := &Object{
		SHA:  want_sha,
		Body: want_body,
		Type: want_type,
	}
	got, _ := r.GetObject(want_sha)
	assert(*got, *want, t)
}

func TestWriteObject(t *testing.T) {
	r := Setup("TestWriteObject")
	defer TearDown("TestWriteObject")

	got_sha, _ := r.WriteObject(want_body, want_type)
	got, _ := r.GetObject(want_sha)
	assert(got_sha, want_sha, t)
	assert(got.Body, want_body, t)
}

func TestParseRawObject(t *testing.T) {
	prefix := append([]byte(fmt.Sprintf("%s %d", want_type, len(want_body))), []byte{'\000'}...)
	s := append(prefix, want_body...)
	got_body, got_type := parseRawObject(s)
	assert(got_type, want_type, t)
	assert(got_body, want_body, t)
}
