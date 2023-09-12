package asciifunctions

import (
	"errors"
	"os"
	"strings"
)

const (
	StandardHash   = "ac85e83127e49ec42487f272d9b9db8b"
	ShadowHash     = "a49d5fcb0d5c59b2e77674aa3ab8bbb1"
	ThinkertoyHash = "86d9947457f6a41a18cb98427e314ff8"
)

func ASCII(inputs []string, font string) (string, error) {
	filename := "fonts/" + font + ".txt"
	if GetHash(filename) != StandardHash && GetHash(filename) != ShadowHash && GetHash(filename) != ThinkertoyHash {
		return "", errors.New("1")
	} else {
		chars, err := ReceiveChars(filename)
		if err != nil {
			return "", err
		}
		var res string
		for i := range inputs {
			switch inputs[i] {
			case "\\n":
				res += "\n"
			case "\\t":
				res += "\t"
			default:
				for line := 0; line < 8; line++ {
					for _, r := range inputs[i] {
						res += chars[r][line]
					}
					if line != 7 {
						res += "\n"
					}
				}
			}
		}
		if !strings.HasSuffix(res, "\n") || !strings.HasPrefix(res, "\n") {
			res += "\n"
		}
		return res, nil
	}
}
func ReceiveChars(source string) (map[rune][]string, error) {
	content, err := os.ReadFile(source)
	if err != nil {
		return map[rune][]string{}, err
	}
	tempstr := strings.ReplaceAll(string(content), "\r", "")
	temp := strings.Split(tempstr[1:], "\n")
	charmap := make(map[rune][]string)
	i := 0
	for key := ' '; key <= '~'; key++ {
		charmap[key] = temp[i : i+8]
		i += 9
	}
	return charmap, nil
}
