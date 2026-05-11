package main

import (
	"fmt"
)

func validateInput(s string) (rune, error) {
	for _, r := range s {
		if r < 32 || r > 126 {
			return r, fmt.Errorf("unsuported charater : %c", r)
		}
	}
	return 0, nil
}
