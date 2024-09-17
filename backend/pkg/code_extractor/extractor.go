package extractor

import (
	"bufio"
	"bytes"
	"fmt"
	"regexp"
	"strings"
)

func ExtractImports(code string, deletePackages bool) (string, string) {
	regex := regexp.MustCompile(`(?i)(import\s*\((?:[^)]*)\)|import\s+\S+|require\s*\(\s*['"][^'"]+['"]\s*\)|from\s+\S+\s+import\s+\S+|#include\s*<[^>]+>|#include\s*"[^"]+"|using\s+namespace\s+\w+;|using\s+[^;]+;)`)
	packageRegex := regexp.MustCompile(`(?i)^package\s+\w+;?`)
	var cleanedCode string

	imports := regex.FindAllString(code, -1)

	if deletePackages {
		cleanedCode = packageRegex.ReplaceAllString(code, "")
		cleanedCode = regex.ReplaceAllString(cleanedCode, "")
	} else {
		cleanedCode = regex.ReplaceAllString(code, "")
	}

	importsString := strings.Join(imports, "\n")
	cleanedCode = strings.TrimSpace(cleanedCode)

	return importsString, cleanedCode
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

func ExtractImportsInsideParenthesis(imports string) []string {
	regex := regexp.MustCompile(`import\s*(?:\(([^)]+)\)|"([^"]+)")`)

	var importblock []string

	matches := regex.FindAllStringSubmatch(imports, -1)
	for _, match := range matches {
		if match[1] != "" {
			inside := match[1]
			importsInside := strings.Split(inside, `" "`)
			for i := range importsInside {
				importsInside[i] = strings.Trim(importsInside[i], `"`)
			}
			importblock = append(importblock, importsInside...)
		} else if match[2] != "" {
			importblock = append(importblock, match[2])
		}
	}

	return importblock
}

func FormatToMultipleImports(imports []string) string {
	formatted := "import (\n"
	for _, imp := range imports {
		formatted += "\t\"" + imp + "\"\n"
	}
	formatted += ")"
	return formatted
}
