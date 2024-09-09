package domains

type ILabRoadService interface {
	GetInventoryInformation(programmingID int32) (inventorys *InventoryInformation, err error)
}

type InventoryInformation struct {
	id            int
	name          string
	dockerImage   string
	cmd           []string
	fileExtension string
	templatePath  string
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

func (i *InventoryInformation) GetFileExtension() string {
	return i.fileExtension
}

func (i *InventoryInformation) GetTemplatePath() string {
	return i.templatePath
}

func NewInventoryInformation(name, dockerImage, fileExtension, templatePath string, ID int, cmd []string) *InventoryInformation {
	return &InventoryInformation{
		id:            ID,
		name:          name,
		dockerImage:   dockerImage,
		fileExtension: fileExtension,
		cmd:           cmd,
		templatePath:  templatePath,
	}
}
