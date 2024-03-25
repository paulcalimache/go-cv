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
	Image       string        `yaml:"image"`
	Contact     Contact       `yaml:"contact"`
	Education   []Education   `yaml:"education"`
	Experiences []Experiences `yaml:"experiences"`
	Skills      []string      `yaml:"skills"`
	Hobbies     []string      `yaml:"hobbies"`
	Projects    []Projects    `yaml:"projects"`
}

type Education struct {
	Timerange   string `yaml:"timerange"`
	Title       string `yaml:"title"`
	Institution string `yaml:"institution"`
}

type Experiences struct {
	Timerange   string `yaml:"timerange"`
	Title       string `yaml:"title"`
	Institution string `yaml:"institution"`
	Description string `yaml:"description"`
}

type Projects struct {
	Title       string `yaml:"title"`
	Description string `yaml:"description"`
}

type Contact struct {
	Mail  string `yaml:"mail"`
	Phone string `yaml:"phone"`
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
