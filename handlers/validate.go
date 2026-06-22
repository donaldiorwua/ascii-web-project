package asciiwebproject

import "errors"

func ValidateInput(inputText string) (string, error) {
	if inputText == "" {
		return "", nil
	}
	if inputText == "\n" {
		return "\n", nil
	}
	for _, word := range inputText {
		if word < 32 || word > 126 {
			return "", errors.New("invalid character, only alphabets or digits are allowed")
		}
	}
	return inputText, nil
}
