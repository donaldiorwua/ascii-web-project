package asciiwebproject

import (
	"errors"
	"os"
	"strings"
)

func LoadBanner(Banner string) ([]string, error) {
	if Banner == "" {
    return nil, errors.New("banner not selected")
}
	filename := "banners/" + Banner + ".txt"
	data, err := os.ReadFile(filename)
	if err != nil {
		return nil, err
	}

	if len(data) == 0 {
		return nil, errors.New("empty banner file")
	}

	lines := strings.Split(string(data), "\n")
	if len(lines) < 855 {
		return  nil, errors.New("incomplete banner file")
	}
	return lines, nil
}
