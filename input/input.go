package input

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"strings"
)

func GetUserInput(label string) (string, error) {
	fmt.Printf("%s", label)

	reader := bufio.NewReader(os.Stdin)
	i, err := reader.ReadString('\n')

	if err != nil {
		return "", err
	}

	if strings.TrimSpace(i) == "" {
		return "", errors.New("empty user input")
	}

	return i, nil
}
