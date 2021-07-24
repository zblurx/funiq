package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
)

func isError(err error) bool {
	if err != nil {
		fmt.Println(err.Error())
		return true
	}
	return false
}

func exists(lines []string, newline string) bool {
	for _, line := range lines {
		if line == newline {
			return true
		}
	}
	return false
}

func main() {
	var verboseMode bool
	flag.BoolVar(&verboseMode, "v", false, "output new file content")
	flag.Parse()

	filename := flag.Arg(0)
	var lines []string

	if filename != "" {
		fd, err := os.OpenFile(filename, os.O_RDWR, 0644)

		if isError(err) {
			os.Exit(1)
		}

		sc := bufio.NewScanner(fd)

		for sc.Scan() {
			if !exists(lines, sc.Text()) {
				lines = append(lines, sc.Text())
			}
		}
		defer fd.Close()

		err = fd.Truncate(0)
		if isError(err) {
			os.Exit(1)
		}

		fd.Seek(0, 0)
		for _, line := range lines {
			if verboseMode {
				fmt.Println(line)
			}
			fmt.Fprintln(fd, line)
		}
	}
}
