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

func (s *parserService) getInventory() (inventory []domains.Inventory, err error) {
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

func (s *parserService) GetLabs() (labs []domains.Labs, err error) {
	// Check if the directory exists
	err = s.checkDir("object")
	if err != nil {
		return nil, err
	}

	// Get list of programming languages
	inventory, err := s.getInventory()
	if err != nil {
		return nil, err
	}

	// Loop through each language
	for _, language := range inventory {
		oneLab := domains.Labs{
			Name:        language.Name,
			DockerImage: language.DockerImage,
			IconPath:    language.IconPath,
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

			var lab domains.Lab
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

func (s *parserService) GetRoads() (roads []domains.Road, err error) {
	// Ensure the "object" directory exists
	err = s.checkDir("object")
	if err != nil {
		return nil, err
	}

	// Retrieve the list of programming languages
	inventory, err := s.getInventory()
	if err != nil {
		return nil, err
	}

	// Iterate over each language in the inventory
	for _, language := range inventory {
		road := domains.Road{
			Name:        language.Name,
			DockerImage: language.DockerImage,
			IconPath:    language.IconPath,
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

			var path domains.Path
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
