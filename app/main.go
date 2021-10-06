package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"

	"github.com/fatih/color"
)

var c_err color.Color = *color.New(color.BgRed).Add(color.FgWhite)
var c_found color.Color = *color.New(color.FgGreen)
var c_not_found color.Color = *color.New(color.FgYellow)

func findPattern(file *os.File, pattern string) bool {
	fileReader := bufio.NewReader(file)
	lines := 0
	lineIdx := 0
	for {
		line, err := fileReader.ReadString('\n')
		lineIdx++
		if strings.Contains(line, pattern) {
			fmt.Printf("%s[%s]: %s", color.GreenString(file.Name()), color.GreenString(strconv.Itoa(lineIdx)), strings.ReplaceAll(line, pattern, color.New(color.Underline).Add(color.FgGreen).Sprintf("%s", pattern)))
			lines++
			continue
		}
		if err == io.EOF {
			return lines > 0
		}
	}
}

func main() {
	if len(os.Args) < 3 {
		fmt.Println(color.CyanString("Usage: "), color.HiGreenString("ggrep pattern filename1 filename2 filename3"))

		return
	}
	pattern, filenames := os.Args[1], os.Args[2:]
	for _, filename := range filenames {
		file, err := os.OpenFile(filename, os.O_RDONLY, os.ModePerm)
		if err != nil {
			c_err.Println(err.Error())
			c_err.DisableColor()
			continue
		}
		found := findPattern(file, pattern)
		if !found {
			c_not_found.Printf("File: %s Pattern Not Found!", filename)
		}
	}
}
