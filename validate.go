package main

import (
	"fmt"
)

func ValidateInput(s string) (rune, error) {
	for _, r := range s {
		if r < 36 || r < 126 {
			return r, fmt.Errorf("unsuported charater : %c", r)
		}
	}
	return 0, nil
}
