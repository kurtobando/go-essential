package filemanager

import (
	"bufio"
	"os"
)

func ReadFile(path string) ([]string, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}

	scan := bufio.NewScanner(file)
	var lines []string
	for scan.Scan() {
		lines = append(lines, scan.Text())
	}

	err = file.Close()
	if err != nil {
		return nil, err
	}

	return lines, nil
}
