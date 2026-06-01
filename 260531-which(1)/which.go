package main

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"strings"
	"time"
)

func hasExeExt(name string) bool {
	ext := strings.ToLower(filepath.Ext(name))
	switch ext {
	case ".exe", ".com", ".bat", ".cmd", ".ps1":
		return true
	default:
		return false
	}
}

// findDuplicates returns a map of duplicate strings
// it`s good, but doesn`t used in the below code
func findDuplicates(a []string) map[string]bool {
	seen := make(map[string]bool)
	dups := make(map[string]bool)
	for _, s := range a {
		if seen[s] {
			dups[s] = true
		} else {
			seen[s] = true
		}
	}
	return dups
}

func hasDuplicates(a []string) bool {
	seen := make(map[string]struct{}, len(a))
	for _, s := range a {
		if _, ok := seen[s]; ok {
			return true
		}
		seen[s] = struct{}{}
	}
	return false
}

func uniqueStrings(a []string) []string {
	seen := make(map[string]struct{}, len(a))
	out := make([]string, 0, len(a))
	for _, s := range a {
		if _, ok := seen[s]; ok {
			continue
		}
		seen[s] = struct{}{}
		out = append(out, s)
	}
	return out
}

func main() {
	start := time.Now()

	arguments := os.Args
	if len(arguments) == 1 {
		fmt.Println("Please provide an argument!")
		return
	}
	file := arguments[1]

	path := os.Getenv("PATH")
	pathSplit := filepath.SplitList(path)

	if hasDuplicates(pathSplit) {
		pathSplit = uniqueStrings(pathSplit)
	}

	for _, directory := range pathSplit {
		fullPath := filepath.Join(directory, file)
		// Does it exist?
		fileInfo, err := os.Stat(fullPath)
		if err != nil {
			continue
		}

		mode := fileInfo.Mode()
		// Is it a regular file?
		if !mode.IsRegular() {
			continue
		}

		// Is it executable?
		if hasExeExt(file) {
			// fmt.Println(mode)
			fmt.Println("Executable:", fullPath)
		}
	}

	fmt.Println()
	b, err1 := json.MarshalIndent(pathSplit, "", "  ")
	if err1 != nil {
		log.Fatal(err1)
	}

	var pathSlice []string
	if err := json.Unmarshal(b, &pathSlice); err != nil {
		log.Fatal(err)
	}
	aStr := make([]string, 0, len(pathSlice))
	for _, item := range pathSlice {
		aStr = append(aStr, item)
	}

	for i, s := range aStr {
		// fmt.Printf("%#v\n", s)
		if strings.Contains(s, `\`) {
			aStr[i] = strings.ReplaceAll(s, `\`, `/`)
		}
		fmt.Printf("%q\n", aStr[i])
	}

	fmt.Printf("Elapsed: %s\n", time.Since(start))

}
