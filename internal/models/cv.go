package models

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
