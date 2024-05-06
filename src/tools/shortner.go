package tools

import (
	"fmt"
)

const (
	BASE = 62
)

func Shorten(longString string) (string, error) {
	return "", nil
}

func Expand(shortString string) (string, error) {
	return "", nil
}

func getByteArray(data string) []byte {
	bytes := make([]bytes, 0)
	for _, r := range data {
		bytes = append(bytes, r)
	}
	return bytes
}

func binStreamify(data []rune) string {
	binStream := ""
	for _, r := range data {
		binStream += fmt.Sprintf("%b", r)
	}
	return binStream
}

func validCharset(data []rune) ([]rune, error) {
	//check if the rune falls under 0-9, a-z, A-Z
	for _, r := range data {
		if !(r >= 48 && r <= 57) && !(r >= 65 && r <= 90) && !(r >= 97 && r <= 122) {
			return nil, fmt.Errorf("Invalid character in the input string %s", string(r))
		}
	}
	return data, nil
}
