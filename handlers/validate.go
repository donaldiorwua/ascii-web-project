package asciiwebproject

import "errors"

func ValidateInput(inputText string) (string, error) {
	if inputText == "" {
		return "", errors.New("empty inputs not allowed, enter some text")
	}
	if inputText == "\n" {
		return "\n", errors.New("new line character returns empty output")
	}
	for _, word := range inputText {
		if word < 32 || word > 126 {
			return inputText, errors.New("invalid character, only printable characters are allowed")
		}
	}
	return inputText, nil
}
