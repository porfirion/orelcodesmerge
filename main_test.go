package main

import (
	"bytes"
	"crypto/md5"
	"fmt"
	"io"
	"os"
	"strings"
	"testing"
)

var inpLocal = `1
2
3
4
5
8
9
`
var inpServer = `5
6
7
8
9
10
`

var outRes = `1
2
3
4
5
6
7
8
9
10
`
var outMis = `6
7
10
`

func Test_merge(t *testing.T) {
	var rdLocal = strings.NewReader(inpLocal)
	var rdServer = strings.NewReader(inpServer)

	var wrRes = bytes.NewBuffer(nil)
	var wrMis = bytes.NewBuffer(nil)

	t.Run("main_test", func(tt *testing.T) {
		merge(rdLocal, rdServer, wrRes, wrMis)
		var res = wrRes.String()
		var mis = wrMis.String()
		if res != outRes {
			tt.Fatalf("wrong result: wanted %s got %s", outRes, res)
		}

		if mis != outMis {
			tt.Fatalf("wrong missing: wanted %s got %s", outMis, mis)
		}

	})
}

func Test_fromFiles(t *testing.T) {
	generate(100000, 50)

	fromFiles()

	hash := func(filename string) string {
		f, err := os.Open(filename)
		if err != nil {
			t.Fatal(err)
		}
		defer f.Close()

		h := md5.New()
		if _, err := io.Copy(h, f); err != nil {
			t.Fatal(err)
		}

		return fmt.Sprintf("%x", h.Sum(nil))
	}

	if hash("output_result.txt") != hash("output_result.txt.example") {
		t.Fatal("wrong result", hash("output_result.txt"), hash("output_result.txt.example"))
	}
	if hash("output_missing.txt") != hash("output_missing.txt.example") {
		t.Fatal("wrong missing", hash("output_missing.txt"), hash("output_missing.txt.example"))
	}
}
