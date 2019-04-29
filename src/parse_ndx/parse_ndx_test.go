package main

import (
    "testing"
)

func TestLoadFile(t *testing.T) {
    filename := "ICRP-07.NDX"
    file := LoadFile(filename)
    file.Close()
}

func TestParseNdx(t *testing.T) {

    if false {
       t.Errorf("error")
    }
}
