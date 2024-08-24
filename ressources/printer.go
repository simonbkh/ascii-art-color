package color

import (
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
)

var (
	Black   = "\x1b[30m"
	Red     = "\x1b[31m"
	Green   = "\x1b[32m"
	Yellow  = "\x1b[33m"
	Blue    = "\x1b[34m"
	Magenta = "\x1b[35m"
	Cyan    = "\x1b[36m"
	White   = "\x1b[37m"
	Reset   = "\x1b[0m"
)

func Printer(inputLine string, slice [][]string, substring string, color string) {
	// Check for non-printable characters
	for j := 0; j < 8; j++ {
		i := 0
		for i < len(inputLine) {
			if (len(os.Args) == 4 || len(os.Args) == 5) && strings.HasPrefix(inputLine[i:], substring) {
				for k := 0; k < len(substring); k++ {
					char := inputLine[i+k]
					index := int(char) - 32
					fmt.Print(getColor(color) + slice[index][j] + Reset)
				}
				i += len(substring) // skips the chars just processed
			} else {
				char := inputLine[i]
				index := int(char) - 32
				if len(os.Args) == 3 {
					fmt.Print(getColor(color) + slice[index][j])
				} else {
					fmt.Print(slice[index][j])
				}
				i++
			}
		}
		fmt.Println()
	}
}

func getColor(colors string) string {
	switch strings.ToLower(colors) {
	case "black":
		return Black
	case "red":
		return Red
	case "green":
		return Green
	case "yellow":
		return Yellow
	case "blue":
		return Blue
	case "magenta":
		return Magenta
	case "cyan":
		return Cyan
	case "white":
		return White
	default:
		if strings.HasPrefix(colors, "rgb(") && strings.HasSuffix(colors, ")") {
			rgbRegex := regexp.MustCompile(`rgb\((\d+),(\d+),(\d+)\)`)
			match := rgbRegex.FindStringSubmatch(colors)
			if match != nil {
				r, _ := strconv.Atoi(match[1])
				g, _ := strconv.Atoi(match[2])
				b, _ := strconv.Atoi(match[3])
				return fmt.Sprintf("\x1b[38;2;%d;%d;%dm", r, g, b)
			}
		}else if strings.HasPrefix(colors, "#"){
			hexRegex := regexp.MustCompile(`#([0-9a-fA-F]{2})([0-9a-fA-F]{2})([0-9a-fA-F]{2})`)
			match := hexRegex.FindStringSubmatch(colors)
			if match != nil {
				r, _ := strconv.ParseInt(match[1], 16, 8)
				g, _ := strconv.ParseInt(match[2], 16, 8)
				b, _ := strconv.ParseInt(match[3], 16, 8)
				return fmt.Sprintf("\x1b[38;2;%d;%d;%dm", r, g, b)
			}
		}
		return ""
	}
}
