package main

import (
	"bytes"
	"io/ioutil"
	"path/filepath"
	"testing"
)

func TestFiles_Config(t *testing.T) {

	raw, err := ioutil.ReadFile(filepath.Join("testdata", "config.yaml"))
	if err != nil {
		t.Fatal(err)
	}
	t.Log(len(raw))

	config, err := BuildConfig("testdata", "testdata")
	if err != nil {
		t.Fatal(err)
	}

	ans, err := config.Marshal()
	if err != nil {
		t.Fatal(err)
	}

	if !bytes.Equal(raw, ans) {
		t.Errorf("byte mismatch:\n \"%s\"\n \"%s\"", string(raw), string(ans))
	}
}
