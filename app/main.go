package main

import (
	"bufio"
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"strconv"
	"strings"

	"github.com/fatih/color"
)

var c_err color.Color = *color.New(color.BgRed).Add(color.FgWhite)
var c_found color.Color = *color.New(color.FgGreen)
var c_not_found color.Color = *color.New(color.FgYellow)

var DELIMETER = strings.Repeat("─", 100)

func findPattern(file *os.File, pattern string) bool {
	fileReader := bufio.NewReader(file)
	lines := 0
	lineIdx := 0

	for {
		line, err := fileReader.ReadString('\n')
		lineIdx++
		if lineIdx == 1 {
			fmt.Println(DELIMETER)
		}
		if strings.Contains(line, pattern) {
			founded := fmt.Sprintf("%s[%s]: %s", color.GreenString(file.Name()), color.GreenString(strconv.Itoa(lineIdx)), strings.ReplaceAll(line, pattern, color.New(color.Underline).Add(color.FgGreen).Sprintf("%s", pattern)))
			fmt.Println(founded + DELIMETER)

			lines++
			continue
		}
		if err == io.EOF {
			return lines > 0
		}
	}
}

func find(pattern string, filenames []string) {
	for _, filename := range filenames {
		file, err := os.OpenFile(filename, os.O_RDONLY, os.ModePerm)
		if err != nil {
			c_err.Println(err.Error())
			c_err.DisableColor()
			continue
		}
		color.Cyan(filename)
		found := findPattern(file, pattern)
		if !found {
			c_not_found.Printf("File: %s Pattern Not Found!\n", filename)
			fmt.Println(DELIMETER)
		}
	}
}

func findInDir(pattern string, dirnames []string) {
	for _, dirname := range dirnames {
		files, err := ioutil.ReadDir(dirname)
		if err != nil {
			c_err.Println(err.Error())
			c_err.DisableColor()
			continue
		}
		for _, file := range files {
			if file.IsDir() {
				findInDir(pattern, []string{dirname + "/" + file.Name()})
				continue
			}
			filePath := dirname + "/" + file.Name()
			find(pattern, []string{filePath})
		}
	}

}

func main() {
	if len(os.Args) < 3 {
		fmt.Println(color.CyanString("Usage: "), color.HiGreenString("ggrep [d] pattern filename1 filename2 filename3"))
		fmt.Println(color.CyanString("d"), color.HiGreenString(" - find in directories"))
		return
	}
	isDir := os.Args[1] == "d"
	if isDir {
		findInDir(os.Args[2], os.Args[3:])
		fmt.Println(isDir)
	} else {
		find(os.Args[1], os.Args[2:])
	}
}
