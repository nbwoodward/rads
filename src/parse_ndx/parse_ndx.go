package main
import (
    "bufio"
    "fmt"
    "log"
    "os"
    "strings"
    "strconv"
)

type Rad struct {
    name string
    hl string
    hls float64
    d1 string
    d1f float64
    d2 string
    d2f float64
    d3 string
    d3f float64
}

func LoadFile(filename string) *os.File {

    fmt.Println("Opening " + filename)
    file, err := os.Open(filename)
    if err != nil {
        log.Fatal(err)
    }

    return file
}



//  Given a string like "56d" meaning 56 days, returns the number of seconds represented
//  by that string.
//  Accepts:
//    miliseconds:    ms
//    seconds:        s
//    minutes:        m
//    hours:          h
//    days:           d
//    years:          y
func ParseHalflife(hl string) float64 {

    last := len(hl) - 1

    var unit string
    var val float64
    var err error

    for last > 0 {
        unit = string(hl[last:])
        val,err = strconv.ParseFloat(hl[:last],64)
        if err == nil {
            break
        }

        last--
    }

    var val_sec float64

    switch unit {
        case "ms":
            val_sec = val / 1000
        case "s":
            val_sec = val
        case "m":
            val_sec = val * 60
        case "h":
            val_sec = val * 60 * 60
        case "d":
            val_sec = val * 60 * 60 * 24
        case "y":
            val_sec = val * 60 * 60 * 24 * 365.2422
    }

    return val_sec
}

//    Given a array of a line from ICRP-97.NDX this parses the daughters.
//    There can be between 1 and 3 daughters and not all lines have them in the
//    same array key.
//    This returns the daughters and branching fractions:
//
//        d1, d2, d3, d1f, d2f, d3f
//
//    Empty daughters are returned as "" with brancing fraction 0.0.
func ParseDaughters(r []string) (string, string, string, float64, float64, float64) {
    var d1, d2, d3 string
    var d1f, d2f, d3f float64
    var err error

    d1_idx := 7
    if strings.Contains(r[3],"SF") {
        d1_idx++
    }

    d1 = r[d1_idx]
    d1f ,err = strconv.ParseFloat(r[d1_idx+2],64)
    if err != nil {
        log.Fatal(err)
    }

    // Rads without a second daughter have r[10] = "0" instead of a rad name
    if r[d1_idx+3] != "0" {
        d2_idx := d1_idx + 3
        d2 = r[d2_idx]
        d2f ,err = strconv.ParseFloat(r[d2_idx+2],64)
        if err != nil {
            log.Fatal(err)
        }

        d3_idx := d2_idx + 3
        if r[d3_idx] != "0" {
            d3 = r[d3_idx]
            d3f ,err = strconv.ParseFloat(r[d3_idx+2],64)
            if err != nil {
                log.Fatal(err)
            } }
    }

    return d1, d2, d3, d1f, d2f, d3f
}

//    Parses a line from ICRP-07.NDX into a Rad object using the utility
//    functions above.
func ParseLine(line string) Rad {

    r := strings.Fields(line)

    radname := r[0]
    hl := r[1]
    hls := ParseHalflife(hl)
    d1, d2, d3, d1f, d2f, d3f := ParseDaughters(r)

    rad := Rad{radname,hl,hls,d1, d1f, d2, d2f, d3,d3f}

    return rad
}

//    Parses an NDX file (specifically ICRP-07.NDX) into a slice of Rad objects.
func ParseNdx( filename string) []Rad {

    file := LoadFile(filename)
    defer file.Close()

    scanner := bufio.NewScanner(file)
    if err := scanner.Err(); err != nil {
        log.Fatal(err)
    }

    var rads []Rad
    count := 0
    for scanner.Scan() {
        if count != 0 {
            rad := ParseLine(scanner.Text())
            rads = append(rads,rad)
        }
        count++
    }

    return rads
}

func main() {
    filename := "ICRP-07.NDX"
    rads  := ParseNdx(filename)
    for _,rad := range rads {
        fmt.Println(rad.name, rad.hl, rad.hls, rad.d1, rad.d1f, rad.d2, rad.d2f, rad.d3, rad.d3f)
    }
}
