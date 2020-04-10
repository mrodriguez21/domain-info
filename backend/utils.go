package main

import (
	"bufio"
	"os/exec"
	"regexp"
	"strings"
)

func StringToLines(s string) (lines []string, err error) {
	scanner := bufio.NewScanner(strings.NewReader(s))
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	err = scanner.Err()
	return
}

func SplitBySpace(s string, n int) (words []string) {
	space := regexp.MustCompile(`\s+`)
	words = space.Split(s, n)
	return
}

func RunCommand(command string, ip string) (output string, err error) {
	cmd := exec.Command(command, ip)
	stdout, err := cmd.Output()

	if err == nil {
		output = string(stdout)
	}
	return
}
