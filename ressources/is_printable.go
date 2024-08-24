package color

import "fmt"

func IsPrintable(sl []string) bool {
	for _, line := range sl {
		for _, char := range line {
			if char < 32 || char > 126 {
				fmt.Printf("Character '%c' is not a printable ASCII character\n", char)
				return false
			}
		}
	}
	return true
}
