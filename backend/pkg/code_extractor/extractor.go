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

	if startIndex != -1 {
		lines = lines[:startIndex] //deletes main from code
	}

	newCode := strings.Join(lines, "\n")
	return newCode, nil
}

func ExtractFuncName(code, newFuncName string) string {
	lines := strings.Split(code, "\n")

	mainFuncRegex := regexp.MustCompile(`(?i)\bmain\b\s*(\(|{|$)`)

	for i, line := range lines {
		if mainFuncRegex.MatchString(line) {
			lines[i] = strings.Replace(line, "main", newFuncName, 1)
			break
		}
	}

	newCode := strings.Join(lines, "\n")
	return newCode
}

// Extracts single and multi typed imports in Golang
func ExtractImportsForGolang(importBlock string) []string {
	re := regexp.MustCompile(`"([^"]+)"`)
	matches := re.FindAllStringSubmatch(importBlock, -1)

	var imports []string
	for _, match := range matches {
		imports = append(imports, match[1])
	}
	return imports
}

// Adds newImports to frontImports in Golang in proper format
func CombineImportsForGolang(existingImports, newImports string) string {
	existingImported := ExtractImportsForGolang(existingImports)
	newImported := ExtractImportsForGolang(newImports)

	uniqueImports := make(map[string]bool)
	var finalList []string

	for _, imp := range existingImported {
		finalList = append(finalList, imp)
		uniqueImports[imp] = true
	}

	for _, imp := range newImported {
		if !uniqueImports[imp] { //Does not add if already exists in existingImports
			finalList = append(finalList, imp)
			uniqueImports[imp] = true
		}
	}

	if len(finalList) == 1 {
		return fmt.Sprintf(`import "%s"`, finalList[0])
	}

	var sb strings.Builder
	sb.WriteString("import (\n")
	for _, imp := range finalList {
		sb.WriteString(fmt.Sprintf("\t\"%s\"\n", imp))
	}
	sb.WriteString(")")

	return sb.String()
}
