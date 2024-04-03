package curriculum

import (
	"errors"
	"os"
	"path"

	"github.com/paulcalimache/go-curriculum/internal/models"
	"gopkg.in/yaml.v3"
)

func ParseFile(filePath string) (*models.CV, error) {
	buf, err := os.ReadFile(filePath)
	if err != nil {
		return nil, err
	}

	switch path.Ext(filePath) {
	case ".yaml", ".yml":
		return parseYamlFile(buf)
	case ".json":
		//TODO
	}
	return nil, errors.New(path.Ext(filePath) + " is not a valid file extension")
}

func parseYamlFile(buf []byte) (*models.CV, error) {
	var cv models.CV
	err := yaml.Unmarshal(buf, &cv)
	if err != nil {
		return nil, err
	}
	return &cv, nil
}
