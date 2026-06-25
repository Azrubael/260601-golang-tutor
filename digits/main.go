// A program for drawing large numbers  2022-08-05
package main

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
	"strings"
)

var bigDigits = [][]string{{
	" 000 ",
	"0   0",
	"0  00",
	"0 0 0",
	"00  0",
	"0   0",
	" 000 "}, {

	"  1  ",
	" 11  ",
	"  1  ",
	"  1  ",
	"  1  ",
	"  1  ",
	" 111 "}, {

	" 222 ",
	"2   2",
	"   2 ",
	"  2  ",
	" 2   ",
	"2    ",
	"22222"}, {

	" 333 ",
	"3   3",
	"    3",
	"   3 ",
	"    3",
	"3   3",
	" 333 "}, {

	"4   4",
	"4   4",
	"4   4",
	"44444",
	"    4",
	"    4",
	"    4"}, {

	"55555",
	"5    ",
	"5    ",
	"5555 ",
	"    5",
	"5   5",
	" 555 "}, {

	" 666 ",
	"6    ",
	"6    ",
	"6666 ",
	"6   6",
	"6   6",
	" 666 "}, {

	"77777",
	"    7",
	"   7 ",
	"  7  ",
	" 7   ",
	"7    ",
	"7    "}, {

	" 888 ",
	"8   8",
	"8   8",
	" 888 ",
	"8   8",
	"8   8",
	" 888 "}, {

	" 999 ",
	"9   9",
	"9   9",
	" 9999",
	"    9",
	"    9",
	" 999 "}}

// The body of "main" procedure
func main() {
	if len(os.Args) == 1 {
		fmt.Printf("usage: %s <whole-number>\n", filepath.Base(os.Args[0]))
		os.Exit(1)
	}

	stringOfDigits := os.Args[1]
	for row := range bigDigits[0] {
		var line strings.Builder
		for column := range stringOfDigits {
			digit := stringOfDigits[column] - '0'
			if 0 < digit && digit <= 9 {
				line.WriteString(bigDigits[digit][row])
				line.WriteString("  ")
			} else {
				log.Fatal("Invalid whole number")
			}
		}
		fmt.Println(line.String())
	}
}
