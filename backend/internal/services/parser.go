package services

import (
	"encoding/json"
	"os"
	"path/filepath"
	"strings"

	"github.com/Yavuzlar/CodinLab/internal/domains"
)

type parserService struct {
	utils IUtilService
}

func newParserService(
	utils IUtilService,
) domains.IParserService {
	return &parserService{
		utils: utils,
	}
}

func (s *parserService) checkDir(dir string) (err error) {
	fileInfo, err := os.Stat(dir)
	if os.IsNotExist(err) {
		return err
	}

	if !fileInfo.IsDir() {
		return os.ErrNotExist
	}

	return nil
}

// Gets json files.
func (s *parserService) findJSONFiles(rootDir string) (jsonFiles []string, err error) {
	// Walk through the directory to find JSON files
	err = filepath.Walk(rootDir, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		// Check if the file is a JSON file
		if !info.IsDir() && strings.HasSuffix(strings.ToLower(info.Name()), ".json") {
			jsonFiles = append(jsonFiles, path)
		}
		return nil
	})

	if err != nil {
		return nil, err
	}
	return
}

func (s *parserService) GetInventory() (inventory []domains.InventoryP, err error) {
	// Check if the directory exists
	err = s.checkDir("object")
	if err != nil {
		return
	}

	// Read the JSON file containing language information
	jsonData, err := os.ReadFile("object/inventory.json")
	if err != nil {
		return nil, err
	}

	// Unmarshal the JSON data into languages slice
	err = json.Unmarshal(jsonData, &inventory)
	if err != nil {
		return nil, err
	}

	return
}

func (s *parserService) GetLabs() (labs []domains.LabsP, err error) {
	// Check if the directory exists
	err = s.checkDir("object")
	if err != nil {
		return nil, err
	}

	// Get list of programming languages
	inventory, err := s.GetInventory()
	if err != nil {
		return nil, err
	}

	// Loop through each language
	for _, language := range inventory {
		oneLab := domains.LabsP{
			ID:          language.ID,
			Name:        language.Name,
			DockerImage: language.DockerImage,
			IconPath:    language.IconPath,
			Cmd:         language.Cmd,
		}

		// Find JSON files for the lab
		jsonFiles, err := s.findJSONFiles(language.LabDir)
		if err != nil {
			return nil, err
		}

		// Loop through each JSON file
		for _, file := range jsonFiles {
			// Read the JSON file
			jsonData, err := os.ReadFile(file)
			if err != nil {
				return nil, err
			}

			var lab domains.LabP
			err = json.Unmarshal(jsonData, &lab)
			if err != nil {
				return nil, err
			}
			oneLab.Labs = append(oneLab.Labs, lab)
			// Append the quest to the lab
		}

		// Append the lab to the labs slice
		labs = append(labs, oneLab)
	}

	return
}

func (s *parserService) GetRoads() (roads []domains.RoadP, err error) {
	// Ensure the "object" directory exists
	err = s.checkDir("object")
	if err != nil {
		return nil, err
	}

	// Retrieve the list of programming languages
	inventory, err := s.GetInventory()
	if err != nil {
		return nil, err
	}

	// Iterate over each language in the inventory
	for _, language := range inventory {
		road := domains.RoadP{
			ID:          language.ID,
			Name:        language.Name,
			DockerImage: language.DockerImage,
			IconPath:    language.IconPath,
			Cmd:         language.Cmd,
		}

		// Locate JSON files within the language's lab directory
		jsonFiles, err := s.findJSONFiles(language.PathDir)
		if err != nil {
			return nil, err
		}

		// Iterate over each JSON file found
		for _, file := range jsonFiles {
			// Read the content of the JSON file
			jsonData, err := os.ReadFile(file)
			if err != nil {
				return nil, err
			}

			var path domains.PathP
			// Unmarshal the JSON data into the path object
			err = json.Unmarshal(jsonData, &path)
			if err != nil {
				return nil, err
			}
			// Append the path to the current road
			road.Paths = append(road.Paths, path)
		}

		// Append the current road to the list of roads
		roads = append(roads, road)
	}
	return
}

func (s *parserService) GetLevels() (levels []domains.LevelP, err error) {
	// Ensure the "object" directory exists
	err = s.checkDir("object")
	if err != nil {
		return nil, err
	}
	// Read the JSON file containing level information
	jsonData, err := os.ReadFile("object/level.json")
	if err != nil {
		return nil, err
	}

	// Unmarshal the JSON data into level slice
	err = json.Unmarshal(jsonData, &levels)
	if err != nil {
		return nil, err
	}
	return
}

func (s *parserService) GetWelcomeBanner() (content []domains.WelcomeContent, err error) {
	// Ensure the "object" directory exists
	err = s.checkDir("object")
	if err != nil {
		return nil, err
	}
	// Read the JSON file containing welcome information
	jsonData, err := os.ReadFile("object/home/welcome.json")
	if err != nil {
		return nil, err
	}

	// Unmarshal the JSON data into level slice
	err = json.Unmarshal(jsonData, &content)
	if err != nil {
		return nil, err
	}
	return
}

func (s *parserService) GetLabBanner() (content []domains.LabContent, err error) {
	// Ensure the "object" directory exists
	err = s.checkDir("object")
	if err != nil {
		return nil, err
	}
	// Read the JSON file containing welcome information
	jsonData, err := os.ReadFile("object/home/lab.json")
	if err != nil {
		return nil, err
	}

	// Unmarshal the JSON data into level slice
	err = json.Unmarshal(jsonData, &content)
	if err != nil {
		return nil, err
	}
	return
}

func (s *parserService) GetRoadBanner() (content []domains.RoadContent, err error) {
	// Ensure the "object" directory exists
	err = s.checkDir("object")
	if err != nil {
		return nil, err
	}
	// Read the JSON file containing welcome information
	jsonData, err := os.ReadFile("object/home/road.json")
	if err != nil {
		return nil, err
	}

	// Unmarshal the JSON data into level slice
	err = json.Unmarshal(jsonData, &content)
	if err != nil {
		return nil, err
	}
	return
}
