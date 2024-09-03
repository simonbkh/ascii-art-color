package main

import (
	"fmt"
	"os"
	"strings"

	color "color/ressources"
)

var (
	colors     string
	outputFile string
)

func formatError() {
	fmt.Println("Usage: go run . [OPTION] [STRING]")
	fmt.Println()
	fmt.Println("EX: go run . --color=<color> <substring to be colored> \"something\"")
}

func optionFlag() bool {
	option := os.Args[1:]

	if strings.HasPrefix(string(option[0]), "--color=") && !strings.HasPrefix(string(option[1]), "--output=") {
		colors = strings.TrimPrefix(string(option[0]), "--color=")
	} else if strings.HasPrefix(string(option[0]), "--output=") && !strings.HasPrefix(string(option[1]), "--color=") {
		outputFile = strings.TrimPrefix(string(option[0]), "--output=")
		color.IsOutput = true
	} else {
		if !strings.HasPrefix(string(option[0]), "--color=") && !strings.HasPrefix(string(option[0]), "--output=") {
			if len(os.Args) > 3 {
				formatError()
				return false
			}
		} else {
			fmt.Println("ss")
			return false
		}
	}

	if color.GetColor(colors) == "" && strings.HasPrefix(option[0], "--color=") {
		fmt.Println("ERROR: the color isn't available.\nRetry with rgb.")
		fmt.Println()
		return false
	}
	return true
}

func main() {
	var banner string
	var input string
	var substring string
	var s string

	if len(os.Args) == 2 {
		if strings.HasPrefix(os.Args[1], "--color=") || strings.HasPrefix(os.Args[1], "--output=") {
			formatError()
			return
		} else {
			input = os.Args[1]
			banner = "standard"
		}
	} else if len(os.Args) == 3 {
		if !optionFlag() {
			return
		}
		if strings.HasPrefix(os.Args[1], "--color=")  || strings.HasPrefix(os.Args[1], "--output="){
			input = os.Args[2]
			banner = "standard"
			
		} else {
			input = os.Args[1]
			banner = os.Args[2]
		}
	} else if len(os.Args) == 4 {
		if !optionFlag() {
			return
		}
		if strings.HasPrefix(os.Args[1], "--color=") {
			substring = os.Args[2]
			input = os.Args[3]
			banner = "standard"
			if !strings.Contains(input, substring) && (input == "standard" || input == "shadow" || input == "thinkertoy") {
				input = os.Args[2]
				banner = os.Args[3]
			}

		}else if color.IsOutput{
			input = os.Args[2]
			banner = os.Args[3]
		}else {
			formatError()
			return
		}
	} else if len(os.Args) == 5 {
		if !optionFlag() {
			return
		}
		if strings.HasPrefix(os.Args[1], "--color=") {
			substring = os.Args[2]
			input = os.Args[3]
			banner = os.Args[4]
			
		} else {
			formatError()
			return
		}
	} else {
		formatError()
		return
	}
	if banner != "standard" && banner != "shadow" && banner != "thinkertoy" {
		fmt.Println("Error: Not a valid banner")
		return
	}

	if !color.IsPrintable(substring) || color.HasWhiteSpace(substring) {
		fmt.Println("\n[Check README file]")
		return
	}
	inputLines := strings.Split(input, "\\n")
	for _, line := range inputLines {
		if !color.IsPrintable(line) {
			return
		}
	}
	str := color.Reader(banner)
	color.Splitter(banner, str)
	if input == "" {
		return
	} else {
		inputLines = color.OnlyNewLine(inputLines)
		for _, value := range inputLines {
			if value == "" {
				if color.IsOutput {
					s += "\n"
				} else {
					fmt.Println()
				}
			} else {
				if color.IsOutput {
					s += color.Printer(value, color.Slice, substring, colors)
				}
				color.Printer(value, color.Slice, substring, colors)
			}
		}
	}
}
