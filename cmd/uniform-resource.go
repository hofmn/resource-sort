package main

import (
	"fmt"
	"os"
	"slices"
	"strings"
)

const (
	LF   = "\n"
	CRLF = "\r\n"
)

var lineSeparator string
var filePaths []string

func main() {
	setLineSeparator()
	setFilePaths()
	for _, path := range filePaths {
		sortAndFormatFile(path)
	}
}

func setLineSeparator() {
	switch strings.ToLower(os.Args[1]) {
	case "--lf":
		lineSeparator = LF
	case "--crlf":
		lineSeparator = CRLF
	default:
		lineSeparator = LF
	}
}

func setFilePaths() {
	if strings.Contains(os.Args[1], "--") {
		filePaths = os.Args[2:]
	} else {
		filePaths = os.Args[1:]
	}
}

func sortAndFormatFile(filePath string) {
	content := mustReadFile(filePath)
	sortedLines := sortFileByLines(content)
	longestKeyLength := getLongestKeyLength(sortedLines)
	formatedLines := formatLines(sortedLines, longestKeyLength)
	if isLineEmpty(formatedLines[0]) {
		formatedLines = formatedLines[1:]
	}
	mustWriteFile(filePath, formatedLines)
}

func mustReadFile(filePath string) string {
	bytes, err := os.ReadFile(filePath)
	if err != nil {
		panic(err)
	}
	return string(bytes)
}

func sortFileByLines(fileContent string) []string {
	lines := getFileLines(fileContent)
	slices.Sort(lines)
	return lines
}

func getFileLines(fileContent string) []string {
	return strings.SplitAfter(fileContent, "\n")
}

func getLongestKeyLength(lines []string) int {
	var longestKeyLength int
	for _, line := range lines {
		lenKey := len(getKey(line))
		if longestKeyLength < lenKey {
			longestKeyLength = lenKey
		}
	}
	return longestKeyLength
}

func formatLines(lines []string, longestKeyLength int) []string {
	var newLines []string
	for i, line := range lines {
		if isLineEmpty(line) {
			continue
		}
		newLine := formatLine(line, longestKeyLength)
		if compareFirstWordFromKeys(getKey(lines[i-1]), getKey(lines[i])) {
			newLines = append(newLines, newLine)
		} else {
			newLines = append(newLines, lineSeparator)
			newLines = append(newLines, newLine)
		}
	}
	return newLines
}

func isLineEmpty(line string) bool {
	return strings.TrimSpace(line) == ""
}

func formatLine(line string, longestKeyLength int) string {
	key := addSpacesBehindKey(getKey(line), longestKeyLength)
	value := replaceSpecialChars(getVal(line))
	newLine := []string{key, "= ", value, lineSeparator}
	return strings.Join(newLine, "")
}

func addSpacesBehindKey(key string, longestKeyLength int) string {
	numberSpacesToAdd := longestKeyLength - len(key)
	spacesToAdd := strings.Repeat(" ", numberSpacesToAdd)
	return strings.Join([]string{key, spacesToAdd}, " ")
}

func getKey(line string) string {
	return strings.TrimSpace(strings.Split(line, "=")[0])
}

func replaceSpecialChars(val string) string {
	var result strings.Builder
	result.Grow(len(val) * 2) // Pre-allocate space for efficiency
	for _, r := range val {
		// Check if the character is in the ASCII range
		if r <= 127 {
			result.WriteRune(r)
			continue
		}
		// For any non-ASCII character, convert to Unicode escape sequence
		result.WriteString(fmt.Sprintf("\\u%04X", r))
	}
	return result.String()
}

func getVal(line string) string {
	return strings.TrimSpace(strings.SplitAfter(line, "=")[1])
}

func compareFirstWordFromKeys(key1 string, key2 string) bool {
	word1 := strings.Split(key1, "_")
	word2 := strings.Split(key2, "_")
	return word1[0] == word2[0]
}

func mustWriteFile(filePathOutput string, content []string) {
	file, err := os.Create(filePathOutput)
	if err != nil {
		panic(err)
	}
	defer file.Close()
	for _, s := range content {
		file.WriteString(s)
	}
}
