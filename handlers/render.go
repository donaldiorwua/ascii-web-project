package asciiwebproject

import (
	"errors"
	"strings"
)

func RenderArt(inputText string, lines []string) (string, error) {
	var Result strings.Builder

	validated, err := ValidateInput(inputText)
	if err != nil {
		return "", errors.New(err.Error())
	}
	splited := strings.Split(validated, `\n`)

	for index, word := range splited {
		if word == "" {
			if index < len(splited)-1 {
				Result.WriteString("\n")
			}
			continue
		}
		for i := range 8 {
			for _, ch := range word {
				start := (int(ch)-32)*9 + 1
				block := lines[start : start+8]
				Result.WriteString(block[i])
			}
			Result.WriteString("\n")
		}
	}
	return Result.String(), nil
}
