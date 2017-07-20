//
// object.go
// Copyright (C) 2017 Akash Kothawale <akash@decached.com>
//
// Distributed under terms of the MIT license.
//

package git

import (
	"bytes"
	"compress/zlib"
	"fmt"
	"io/ioutil"
	"os"
)

type Object struct {
	SHA  string
	Body []byte
	Type string
}

func GetObject(sha string) (*Object, error) {
	r, err := os.Open(fmt.Sprintf("%s/objects/%s/%s", home, sha[:2], sha[2:]))
	if err != nil {
		return nil, err
	}

	z, err := zlib.NewReader(r)
	if err != nil {
		return nil, err
	}
	defer z.Close()

	raw, err := ioutil.ReadAll(z)
	body, type_ := parseRawObject(raw)
	if err != nil {
		return nil, err
	}

	return &Object{SHA: sha, Body: body, Type: type_}, nil
}

func WriteObject(r []byte, type_ string) (string, error) {
	prefix := append([]byte(fmt.Sprintf("%s %d", type_, len(r))), []byte{'\000'}...)
	raw := append(prefix, r...)
	sha := hash(raw)

	objD := fmt.Sprintf("%s/objects/%s", home, sha[:2])
	if _, err := os.Stat(objD); os.IsNotExist(err) {
		os.Mkdir(objD, os.ModePerm)
	}

	objF := fmt.Sprintf("%s/objects/%s/%s", home, sha[:2], sha[2:])
	if _, err := os.Stat(objF); os.IsNotExist(err) {
		wF, err := os.Create(objF)
		checkErr(err)
		defer wF.Close()

		w := zlib.NewWriter(wF)
		defer w.Close()
		w.Write(raw)
	}

	return sha, nil
}

func parseRawObject(z []byte) ([]byte, string) {
	// Structure of any git object: "type length\0body"
	a := bytes.Split(z, []byte{'\000'})
	b := bytes.Split(a[0], []byte{' '})

	return a[1], string(b[0])
}
