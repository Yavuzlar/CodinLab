package domains

type ICommonService interface {
	GetInventoryInformation(programmingID, language string) (inventories *InventoryInformation, err error)
}

type InventoryInformation struct {
	id            int
	name          string
	dockerImage   string
	cmd           []string
	shCmd         []string
	pathDir       string
	fileExtension string
	templatePath  string
	language      InventoryLanguage
}

func (i *InventoryInformation) GetID() int {
	return i.id
}

func (i *InventoryInformation) GetName() string {
	return i.name
}

func (i *InventoryInformation) GetDockerImage() string {
	return i.dockerImage
}

func (i *InventoryInformation) GetCmd() []string {
	return i.cmd
}

func (i *InventoryInformation) GetShCmd() []string {
	return i.shCmd
}

func (i *InventoryInformation) GetPathDir() string {
	return i.pathDir
}

func (i *InventoryInformation) GetFileExtension() string {
	return i.fileExtension
}

func (i *InventoryInformation) GetTemplatePath() string {
	return i.templatePath
}

func (i *InventoryInformation) GetLanguage() *InventoryLanguage {
	return &i.language
}

func (i *InventoryInformation) SetLanguage(language InventoryLanguage) {
	i.language = language
}

func NewInventoryInformation(name, dockerImage, fileExtension, pathDir string, ID int, cmd, shCmd []string, infoLanguage InventoryLanguage) *InventoryInformation {
	info := &InventoryInformation{
		id:            ID,
		name:          name,
		dockerImage:   dockerImage,
		fileExtension: fileExtension,
		cmd:           cmd,
		shCmd:         shCmd,
		pathDir:       pathDir,
		language:      infoLanguage,
	}

	return info
}

type InventoryLanguage struct {
	lang        string
	title       string
	description string
}

func (il *InventoryLanguage) GetLang() string {
	return il.lang
}

func (il *InventoryLanguage) SetLang(lang string) {
	il.lang = lang
}

func (il *InventoryLanguage) GetTitle() string {
	return il.title
}

func (il *InventoryLanguage) SetTitle(title string) {
	il.title = title
}

func (il *InventoryLanguage) GetDescription() string {
	return il.description
}

func (il *InventoryLanguage) SetDescription(description string) {
	il.description = description
}

func NewInventoryLanguage(lang, title, description string) *InventoryLanguage {
	return &InventoryLanguage{
		lang:        lang,
		title:       title,
		description: description,
	}
}
