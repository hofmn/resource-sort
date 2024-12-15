package main

import (
	"os"
	"slices"
	"strings"
)

const (
	filePathInput  string = "test.txt"
	filePathOutput string = "output.txt"
)

func main() {
	sortAndFormatFile(filePathInput, filePathOutput)
}

func sortAndFormatFile(filePathInput, filePathOutput string) {
	content := mustReadFile(filePathInput)
	sortedLines := sortFileByLines(content)
	longestKeyLength := getLongestKeyLength(sortedLines)
	formatedLines := formatLines(sortedLines, longestKeyLength)
	mustWriteFile(filePathOutput, formatedLines)
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
	value := replaceUmlauts(getVal(line))
	newLine := []string{key, "= ", value, "\n"}
	return strings.Join(newLine, "")
}

func addSpacesBehindKey(key string, longestKeyLength int) string {
	numberSpacesToAdd := longestKeyLength - len(key)
	spacesToAdd := strings.Repeat(" ", numberSpacesToAdd)
	return strings.Join([]string{key, spacesToAdd}, " ")
}

func replaceUmlauts(val string) string {
	replacements := map[string]string{
		"ä": "\\u00E4",
		"Ä": "\\u00C4",
		"ü": "\\u00FC",
		"Ü": "\\u00DC",
		"ö": "\\u00F6",
		"Ö": "\\u00D6",
		"ß": "\\u00DF",
	}
	if strings.ContainsAny(val, "äÄüÜöÖß") {
		for old, new := range replacements {
			val = strings.ReplaceAll(val, old, new)
		}
	}
	return val
}

func getVal(line string) string {
	return strings.TrimSpace(strings.SplitAfter(line, "=")[1])
}
