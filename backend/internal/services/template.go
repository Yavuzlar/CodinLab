package services

import (
	"fmt"
	"strings"

	"github.com/Yavuzlar/CodinLab/internal/domains"
)

type templateService struct {
	utils         IUtilService
	logService    domains.ILogService
	parserService domains.IParserService
	labService    domains.ILabService
	roadService   domains.IRoadService
}

func newTemplateService(
	utils IUtilService,
	logService domains.ILogService,
	parserService domains.IParserService,
	labService domains.ILabService,
	roadService domains.IRoadService,
) domains.ITemplateService {
	return &templateService{
		utils:         utils,
		logService:    logService,
		parserService: parserService,
		labService:    labService,
		roadService:   roadService,
	}
}

// Template Generator ana fonksiyonu
func (s *templateService) TemplateGenerator(category string, programmingID, labRoadID int) (string, error) {
	template := ""
	if category == domains.TypeLab {
		lab, err := s.labService.GetLabsFilter("", programmingID, labRoadID, nil, nil)
		if err != nil {
			return "", err
		}
		if lab[0].GetName() == "GO" {
			selectedLab := lab[0].GetLabs()
			selectedQuest := selectedLab[0].GetQuest()
			template, err = s.goTemplateGenerator(selectedQuest.GetFuncName(), selectedQuest.GetParams(), selectedQuest.GetReturns(), selectedQuest.GetQuestImports())
			if err != nil {
				return "", err
			}
		}

	}

	if category == domains.TypeRoad {
		road, err := s.roadService.GetRoadFilter("", programmingID, labRoadID, nil, nil)
		if err != nil {
			return "", err
		}

		if road[0].GetName() == "GO" {
			selectedRoad := road[0].GetPaths()
			selectedQuest := selectedRoad[0].GetQuest()
			template, err = s.goTemplateGenerator(selectedQuest.GetFuncName(), selectedQuest.GetParams(), selectedQuest.GetReturns(), selectedQuest.GetQuestImports())
			if err != nil {
				return "", err
			}
		}
	}

	return template, nil
}

// GO dili ana fonksiyonu
func (s *templateService) goTemplateGenerator(funcName string, params []domains.Param, returns []domains.Returns, imports []string) (template string, err error) {
	template += "package main \n" //adding main package

	template += "\nimport " //adding imports packages
	template += formatImports(imports)
	template += "\n\nfunc " + funcName + "("

	if len(params) > 0 {
		for i, param := range params {
			template += param.GetName() + " " + param.GetType()
			if len(params) != i+1 {
				template += ","
			}
		}
	}
	template += ") "
	template += "("
	if len(returns) > 0 {

		for i, ret := range returns {
			template += ret.GetName() + " " + ret.GetType()
			if len(params) != i+1 {
				template += ","
			}
		}

	}
	template += ") "
	template += "{\n    //Fill in the function\n}"

	template += "\n\nfunc main() {"

	template += goDeclareParams(params, returns)                                                                      // params ve returns değişkenleri tanımlandı.
	template += "\n" + addTab(1) + goReadFromInput(params)                                                            //params değerleri klavyeden alındı.
	template += "\n" + addTab(1) + formatReturnNames(returns) + "=" + funcName + "(" + formatParamNames(params) + ")" //fonksiyon çağırıldı.

	if strings.Contains(formatReturnNames(returns), "err") { //eğer fonksiyon err döndürüyorsa , err kontrolü yapılır.
		template += "\n" + addTab(1) + "if err != nil {"
		template += "\n" + addTab(1) + "fmt.Println(err) }"
	}

	template += "\n" + addTab(1) + goFormatPrintStatement(returns) //fonksiyonun err hariç tüm değişkenlerini ekrana yazdırıldı.
	template += "\n}"

	return template, nil
}

func addTab(tabCount int) (tab string) {
	for i := 0; i < tabCount; i++ {
		tab += "    "
	}

	return tab
}

// importsları işleyen fonksiyon
func formatImports(imports []string) string {
	if len(imports) == 0 {
		return ""
	}

	if len(imports) == 1 {
		return fmt.Sprintf("\"%s\"", imports[0])
	}

	var formattedImports []string
	for _, imp := range imports {
		formattedImports = append(formattedImports, fmt.Sprintf("\"%s\"", imp))
	}
	return "(\n" + addTab(1) + strings.Join(formattedImports, "\n"+addTab(1)) + "\n)"
}

// Return isimlerini virgüllerle ayrılmış bir string olarak döndüren fonksiyon
func formatReturnNames(returns []domains.Returns) string {
	var names []string
	for _, r := range returns {
		names = append(names, r.GetName())
	}
	return strings.Join(names, ",")
}

// Param isimlerini virgüllerle ayrılmış bir string olarak döndüren fonksiyon
func formatParamNames(params []domains.Param) string {
	var names []string
	for _, p := range params {
		names = append(names, p.GetName())
	}
	return strings.Join(names, ",")
}

// parametre ve returnsların mainde tanımlandığı fonksiyon
func goDeclareParams(params []domains.Param, returns []domains.Returns) string {
	typeMap := make(map[string][]string)

	for _, param := range params {
		typeMap[param.GetType()] = append(typeMap[param.GetType()], param.GetName())
	}

	for _, ret := range returns {
		typeMap[ret.GetType()] = append(typeMap[ret.GetType()], ret.GetName())
	}

	var builder strings.Builder

	for t, names := range typeMap {

		if len(names) > 0 {
			builder.WriteString(fmt.Sprintf("\n"+addTab(1)+"var %s %s", strings.Join(names, ", "), t))
		}
	}

	return builder.String()
}

// Klavyeden array alma ortak stringleri içeren fonksiyon
func goReadArrayCommon(count int) string {
	var returnedString string

	if count == 0 {
		returnedString += "\n" + addTab(1) + "reader := bufio.NewReader(os.Stdin)\n"
		returnedString += addTab(1) + "input, _ := reader.ReadString('\\n')\n"
		returnedString += addTab(1) + "input = strings.TrimSpace(input)\n"
		returnedString += "\n"
		returnedString += addTab(1) + "stringValues := strings.Split(input, \" \")\n"
	} else {
		returnedString += addTab(1) + "input, _ = reader.ReadString('\\n')\n"
		returnedString += addTab(1) + "input = strings.TrimSpace(input)\n"
		returnedString += "\n"
		returnedString += addTab(1) + "stringValues = strings.Split(input, \" \")\n"
	}

	return returnedString
}

// Array türleri için dönüşüm fonksiyonunu döndüren fonksiyon
func goConvertArray(arrayType, conversionFunc, arrayName string) string {
	var returnedString string
	returnedString += addTab(1) + arrayName + "= make(" + arrayType + ", len(stringValues))\n"
	returnedString += addTab(1) + "for i, str := range stringValues {\n"
	returnedString += addTab(2) + "val, err := " + conversionFunc + "(str)\n"
	returnedString += addTab(2) + "if err != nil {\n"
	returnedString += addTab(3) + "fmt.Println(\"Hatalı giriş, lütfen uygun veri türü girin.\")\n"
	returnedString += addTab(3) + "return\n"
	returnedString += addTab(2) + "}\n"
	returnedString += addTab(2) + arrayName + "[i] = val\n"
	returnedString += addTab(1) + "}\n"
	return returnedString
}

// Return isimlerini ve format dizesini döndüren fonksiyon
func goFormatPrintStatement(returns []domains.Returns) string {
	var names []string
	var formatParts []string

	for _, r := range returns {
		if r.GetName() != "err" {
			names = append(names, r.GetName())
			formatParts = append(formatParts, "%v")
		}
	}

	if len(names) == 0 {
		return "Hiçbir değer bulunamadı."
	}

	formatString := fmt.Sprintf("fmt.Printf(\"%s\", %s)", strings.Join(formatParts, " "), strings.Join(names, ", "))
	return formatString
}

// interface{} al ve 1. değer tipi 2. değer değişkenin değeri olacak şekilde parse eden fonksiyon
func goGetInterface(funcName string) string {
	var returnedString string

	returnedString += addTab(1) + "\nfor i := 0; i < len(stringValues); i = i + 2 {\n"
	returnedString += addTab(2) + "if stringValues[i+1] != \"\" {\n"
	returnedString += addTab(3) + "switch stringValues[i] {\n"
	returnedString += addTab(4) + "case \"string\":\n"
	returnedString += addTab(5) + funcName + " = append(" + funcName + ", stringValues[i+1])\n"
	returnedString += addTab(4) + "case \"bool\":\n"
	returnedString += addTab(5) + "parsedValue, err := strconv.ParseBool(stringValues[i+1])\n"
	returnedString += addTab(5) + "if err != nil {\n"
	returnedString += addTab(6) + "fmt.Println(\"Geçersiz bool değeri.\")\n"
	returnedString += addTab(6) + "continue\n"
	returnedString += addTab(5) + "}\n"
	returnedString += addTab(5) + funcName + " = append(" + funcName + ", parsedValue)\n"
	returnedString += addTab(4) + "case \"int\":\n"
	returnedString += addTab(5) + "parsedValue, err := strconv.Atoi(stringValues[i+1])\n"
	returnedString += addTab(5) + "if err != nil {\n"
	returnedString += addTab(6) + "fmt.Println(\"Geçersiz int değeri.\")\n"
	returnedString += addTab(6) + "continue\n"
	returnedString += addTab(5) + "}\n"
	returnedString += addTab(5) + funcName + " = append(" + funcName + ", parsedValue)\n"
	returnedString += addTab(4) + "case \"float\":\n"
	returnedString += addTab(5) + "parsedValue, err := strconv.ParseFloat(stringValues[i+1], 64)\n"
	returnedString += addTab(5) + "if err != nil {\n"
	returnedString += addTab(6) + "fmt.Println(\"Geçersiz float değeri.\")\n"
	returnedString += addTab(6) + "continue\n"
	returnedString += addTab(5) + "}\n"
	returnedString += addTab(5) + funcName + " = append(" + funcName + ", parsedValue)\n"
	returnedString += addTab(3) + "}\n"
	returnedString += addTab(2) + "}\n"
	returnedString += addTab(1) + "}\n"

	return returnedString
}

// Klavyeden değer alma fonksiyonu
func goReadFromInput(params []domains.Param) string {
	var codeSnippet string
	var formatSpecifiers []string
	var variables []string
	codeSnippetCount := 0

	for _, param := range params {
		switch param.GetType() {
		case "int":
			formatSpecifiers = append(formatSpecifiers, "%v")
			variables = append(variables, "&"+param.GetName())
		case "float":
			formatSpecifiers = append(formatSpecifiers, "%v")
			variables = append(variables, "&"+param.GetName())
		case "string":
			formatSpecifiers = append(formatSpecifiers, "%v")
			variables = append(variables, "&"+param.GetName())
		case "bool":
			formatSpecifiers = append(formatSpecifiers, "%v")
			variables = append(variables, "&"+param.GetName())
		case "[]int":
			codeSnippet += goReadArrayCommon(codeSnippetCount) + goConvertArray("[]int", "strconv.Atoi", param.GetName())
			codeSnippetCount++
			continue
		case "[]float":
			codeSnippet += goReadArrayCommon(codeSnippetCount) + goConvertArray("[]float64", "strconv.ParseFloat", param.GetName())
			codeSnippetCount++
			continue
		case "[]string":
			codeSnippet += goReadArrayCommon(codeSnippetCount) + "\n" + addTab(1) + param.GetName() + " = stringValues\n"
			codeSnippetCount++
			continue
		case "[]bool":
			codeSnippet += goReadArrayCommon(codeSnippetCount) + goConvertArray("[]bool", "strconv.ParseBool", param.GetName())
			codeSnippetCount++
			continue
		case "[]interface{}":
			codeSnippet += goReadArrayCommon(codeSnippetCount) + goGetInterface(param.GetName())
			codeSnippetCount++
			continue
		}

	}

	if len(formatSpecifiers) > 0 {
		formatString := strings.Join(formatSpecifiers, " ")
		variableList := strings.Join(variables, ", ")
		codeSnippet += fmt.Sprintf("fmt.Scanf(\"%s\", %s)\n", formatString, variableList)
	}

	return codeSnippet
}
