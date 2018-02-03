//
// helper_test.go
// Copyright (C) 2017 Akash Kothawale <akash@decached.com>
//
// Distributed under terms of the MIT license.
//

package git

import "testing"

func TestHash(t *testing.T) {
	want := "4b4fcadfe74e1d0c88efec01412a98af8920ea3d"
	got := hash([]byte("verify hash"))
	assert(got, want, t)
}
