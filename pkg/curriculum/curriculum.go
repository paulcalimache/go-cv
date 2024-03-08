package curriculum

import (
	"log/slog"
	"os"

	"gopkg.in/yaml.v3"
)

type CV struct {
	Firstname   string        `yaml:"firstname"`
	Lastname    string        `yaml:"lastname"`
	Job         string        `yaml:"job"`
	Description string        `yaml:"description"`
	Education   []Education   `yaml:"education"`
	Experiences []Experiences `yaml:"experiences"`
	Skills      []string      `yaml:"skills"`
}

type Education struct {
	Timerange   string `yaml:"timerange"`
	Title       string `yaml:"title"`
	Institution string `yaml:"institution"`
}

type Experiences struct {
	Timerange   string `yaml:"timerange"`
	Title       string `yaml:"title"`
	Company     string `yaml:"company"`
	Description string `yaml:"description"`
}

func ParseCV(file string) (*CV, error) {
	slog.Info("Parsing file " + file + " ...")
	buf, err := os.ReadFile(file)
	if err != nil {
		return nil, err
	}

	var cv CV

	err = yaml.Unmarshal(buf, &cv)
	if err != nil {
		return nil, err
	}
	slog.Info(cv.Firstname + " CV successfully parsed !")
	return &cv, nil
}
