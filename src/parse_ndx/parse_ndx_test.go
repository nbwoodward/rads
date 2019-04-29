package main

import (
    "testing"
    "fmt"
)

func TestLoadFile(t *testing.T) {
    filename := "ICRP-07.NDX"
    file := LoadFile(filename)
    file.Close()
}

func TestParseHalflife(t *testing.T) {
    tests := []struct {
        hl string
        sec float64
    }{
        {"10ms", .01},
        {"10s", 10},
        {"10m", 600},
        {"10h", 36000},
        {"10d", 864000},
        {"10y", 315569260.8},
        {"7E+15y",2.2089848256e+23},
    }

    for _,r := range tests {
        hls := ParseHalflife(r.hl)
        fmt.Println("Trying", r.hl, "should be",r.sec,"got", hls)

        if hls != r.sec {
           t.Errorf("Halflife was incorrect, entered %s, should have been %f, got %f", r.hl, r.sec, hls)
        }

    }

}

func TestParseDaughters(t *testing.T) {
    /*
        Test:
            1 daughter
            2 daughter
            3 daughter
            line with SF
    */
    if false {
       t.Errorf("error")
    }
}

func TestParseLine(t *testing.T) {
    if false {
       t.Errorf("error")
    }
}


/*
func TestParseNdx(t *testing.T) {
    if false {
       t.Errorf("error")
    }
}
*/
