package ui

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func Confirm(question string, cancel bool) (bool, error) {
	tty, err := os.Open("/dev/tty")
	if err != nil {
		return false, nil
	}
	defer tty.Close()

	reader := bufio.NewReader(tty)

	cancelText := ""
	if cancel {
		cancelText = " or (Ctrt+C to cancel)"
	}

	fmt.Printf("%s (Y/n)%s: ", question, cancelText)

	answer, err := reader.ReadString('\n')
	if err != nil {
		return false, err
	}

	answer = strings.ToLower(strings.TrimSpace(answer))

	return (answer == "Y" || answer == "y" || answer == "yes"), nil
}
