package extractor

import (
	"bufio"
	"bytes"
	"fmt"
	"regexp"
	"strings"
)

func ExtractImports(code string) string {
	regex := regexp.MustCompile(`(?i)(package\s+\w+[\s\S]*?(?:import\s*\((?:[^)]*)\)|import\s+(?:[^;\s]+))|import\s*\([^)]*\)|import\s+\S+|require\s*\(\s*['"][^'"]+['"]\s*\)|from\s+\S+\s+import\s+\S+|#include\s*<[^>]+>|#include\s*"[^"]+"|using\s+namespace\s+\w+;|using\s+[^;]+;)`)

	cleanedCode := regex.ReplaceAllString(code, "")
	cleanedCode = strings.TrimSpace(cleanedCode)

	return cleanedCode
}

func ExtractMainFunction(code string) (string, error) {
	scanner := bufio.NewScanner(bytes.NewReader([]byte(code)))

	var lines []string
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		return "", fmt.Errorf("error while reading code: %w", err)
	}

	mainFuncRegex := regexp.MustCompile(`(?i)\bmain\b\s*(\(|{|$)`)

	startIndex := -1
	for i, line := range lines { //searches for main in code
		if mainFuncRegex.MatchString(line) {
			startIndex = i
			break
		}
	}

	if startIndex == -1 {
		return "", fmt.Errorf("main function not found")
	}
	lines = lines[:startIndex] //deletes main from code

	newCode := strings.Join(lines, "\n")
	return newCode, nil
}
