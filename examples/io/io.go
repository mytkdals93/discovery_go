package io

import (
	"bufio"
	"fmt"
	"io"
)

func WriteTo(w io.Writer, lines []string) error {
	for _, line := range lines {
		_, err := fmt.Fprintln(w, line)
		if err != nil {
			return err
		}
	}
	return nil
}

func ReadFrom(r io.Reader, lines *[]string) error {

	scanner := bufio.NewScanner(r)
	for scanner.Scan() {
		*lines = append(*lines, scanner.Text())
	}
	if err := scanner.Err(); err != nil {
		return err
	}
	return nil
}
