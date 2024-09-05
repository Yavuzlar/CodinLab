package extractor

import (
	"bufio"
	"bytes"
	"regexp"
	"strings"

	service_errors "github.com/Yavuzlar/CodinLab/internal/errors"
)

func ExtractImports(code string) string {
	regex := regexp.MustCompile(`import\s*(?:\(([^)]+)\)|"([^"]+)")`)

	var importblock string

	matches := regex.FindStringSubmatch(code)
	if len(matches) > 0 {
		if matches[1] != "" {
			importblock = matches[1]
		} else if matches[2] != "" {
			importblock = "\"" + matches[2] + "\""
		}
	}

	// Eğer import bloğu varsa ve fmt veya reflect import edilmemişse eklenir
	if !strings.Contains(importblock, "\"fmt\"") {
		importblock += "\n    \"fmt\""
	}
	if !strings.Contains(importblock, "\"reflect\"") {
		importblock += "\n    \"reflect\""
	}

	if importblock != "" {
		return "import(\n" + importblock + "\n)"
	}

	// Eğer hiç import bulunamadıysa fmt ve reflect'i ekle
	return "import(\n    \"fmt\"\n    \"reflect\"\n)"
}

func ExtractFunction(code, funcName string) (string, error) {
	var functionBody strings.Builder
	var inFunction bool
	var braceCount int

	scanner := bufio.NewScanner(bytes.NewReader([]byte(code)))
	for scanner.Scan() {
		line := scanner.Text()
		if strings.Contains(line, "func "+funcName+"(") {
			inFunction = true
		}

		if inFunction {
			functionBody.WriteString(line + "\n")
			braceCount += strings.Count(line, "{") - strings.Count(line, "}")
			if braceCount == 0 {
				break
			}
		}
	}

	if err := scanner.Err(); err != nil {
		return "", service_errors.NewServiceErrorWithMessage(500, "error while reading code")
	}

	if !inFunction {
		return "", service_errors.NewServiceErrorWithMessage(500, "function name is invalid")
	}

	return functionBody.String(), nil
}
