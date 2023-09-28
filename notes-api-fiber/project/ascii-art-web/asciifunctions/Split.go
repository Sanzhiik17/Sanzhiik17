package asciifunctions

import (
	"errors"
	"regexp"
)

func Split(input string) ([]string, error) {
	var new_slice []string
	var word string
	re, err := regexp.Compile(`\n|\\n|\\t|[ -~]`)
	if err == nil && !re.MatchString(input) || err != nil {
		return []string{}, errors.New("ERROR: non-ascii aguments contained\n")
	}
	slice := re.FindAllString(input, -1)
	for _, v := range slice {
		switch {
		case v != "\\n" && v != "\n" && v != "\\t":
			word += v
		case v == "\\t" || v == "\t":
			word += "    "
		case v == "\n":
			if word != "" {
				new_slice = append(new_slice, word)
			}
			new_slice = append(new_slice, "\\n")
			word = ""
		case word != "":
			new_slice = append(new_slice, word)
			new_slice = append(new_slice, v)
			word = ""
		default:
			new_slice = append(new_slice, v)
		}
	}
	if word != "" {
		new_slice = append(new_slice, word)
	}
	return new_slice, nil
}
