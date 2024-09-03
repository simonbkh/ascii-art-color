package color

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

var (
	IsOutput bool
	Black   = "\x1b[38;2;0;0;0m"
	Red     = "\x1b[38;2;255;0;0m"
	Green   = "\x1b[38;2;0;255;0m"
	Yellow  = "\x1b[38;2;255;255;0m"
	Blue    = "\x1b[38;2;0;0;255m"
	Magenta = "\x1b[38;2;0;0;255m"
	Cyan    = "\x1b[38;2;43;255;255m"
	White   = "\x1b[38;2;255;255;255m"
	Orange = "\x1b[38;2;255;140;0m"
	Reset   = "\x1b[0m"
)

// Prints the inputs as a ASCII ART
func Printer(inputLine string, slice [][]string, substring string, color string)string {
	str := ""
	// Check for non-printable characters
	for j := 0; j < 8; j++ {
		i := 0
		for i < len(inputLine) {
			if (len(os.Args) == 4 || len(os.Args) == 5) && strings.HasPrefix(inputLine[i:], substring) {
				for k := 0; k < len(substring); k++ {
					char := inputLine[i+k]
					index := int(char) - 32
					if IsOutput {
						str += slice[index][j] 
					}
					fmt.Print(GetColor(color) + slice[index][j] + Reset)
				}
				i += len(substring) // skips the chars just processed
			} else {
				char := inputLine[i]
				index := int(char) - 32
				if IsOutput {
					str += slice[index][j] 
				}else{
				if len(os.Args) == 3 {
					fmt.Print(GetColor(color) + slice[index][j])
				} else {
					fmt.Print(slice[index][j])
				}
				}
				
				i++
			}
		}
		fmt.Println()
	}
	return str
}

func GetColor(colors string) string {
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
	case "orange":
		return Orange
	default:
		if strings.HasPrefix(colors, "rgb(") && strings.HasSuffix(colors, ")") {
			inner := colors[4 : len(colors)-1]
			components := strings.Split(inner, ", ")
			if len(components) == 3 {
				r, err := strconv.Atoi(components[0])
				g, err1 := strconv.Atoi(components[1])
				b, err2 := strconv.Atoi(components[2])
				
				if (err == nil && err1 == nil && err2 == nil) &&  (r <= 255 && r >=0) && (g <= 255 && g >=0)&&  (b <= 255 && b >=0){
					return "\x1b[38;2;" + components[0] + ";" + components[1] + ";" + components[2] + "m"
				}
			}
			fmt.Println("ERROR: the color isn't available.\nRetry with a valid RGB code.")
			fmt.Println()
			os.Exit(69)
		}
		return ""
	}
}
