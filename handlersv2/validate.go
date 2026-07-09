package asciiwebproject

import (
	"errors"
	"strings"
)

func ValidateInput(inputText string) (string, error) {
	inputText = strings.ReplaceAll(inputText, "\r\n", "\n")
	inputText = strings.ReplaceAll(inputText, "\\n", "\n")

	if inputText == "" {
		return "", errors.New("empty inputs not allowed, enter some text")
	}

	if inputText == "\n" {
		return "", errors.New("new line character returns empty output")
	}
	for _, word := range inputText {
		if word != '\n' && (word < 32 || word > 126) {
			return inputText, errors.New("invalid character, only printable characters are allowed")
		}
	}
	return inputText, nil
}
