package main

import (
	"fmt"
	"os"
	"slices"
	"strings"
)

func main() {
	filePaths := os.Args[1:]
	for _, path := range filePaths {
		sortAndFormatFile(path)
	}
}

func sortAndFormatFile(filePath string) {
	content := mustReadFile(filePath)
	sortedLines := sortFileByLines(content)
	longestKeyLength := getLongestKeyLength(sortedLines)
	formatedLines := formatLines(sortedLines, longestKeyLength)
	mustWriteFile(filePath, formatedLines)
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
	for _, line := range lines {
		if isLineEmpty(line) {
			continue
		}
		newLine := formatLine(line, longestKeyLength)
		newLines = append(newLines, newLine)
	}
	return newLines
}

func getKey(line string) string {
	return strings.TrimSpace(strings.Split(line, "=")[0])
}

func isLineEmpty(line string) bool {
	return strings.TrimSpace(line) == ""
}

func formatLine(line string, longestKeyLength int) string {
	key := addSpacesBehindKey(getKey(line), longestKeyLength)
	value := replaceSpecialChars(getVal(line))
	newLine := []string{key, "= ", value, "\n"}
	return strings.Join(newLine, "")
}

func addSpacesBehindKey(key string, longestKeyLength int) string {
	numberSpacesToAdd := longestKeyLength - len(key)
	spacesToAdd := strings.Repeat(" ", numberSpacesToAdd)
	return strings.Join([]string{key, spacesToAdd}, " ")
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
