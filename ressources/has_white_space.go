package color

import "fmt"

func HasWhiteSpace(substring string) bool {
	slice := []string{`\a`, `\b`, `\t`, `\n`, `\v`, `\f`, `\r`}
	for i := 0; i < len(slice); i++ {
		for j := 0; j < len(substring)-1; j++ {
			if substring[j:j+2] == slice[i] {
				fmt.Println("White spaces NOT allowed in the given substring")
				return true
			}
		}
	}
	return false
}
