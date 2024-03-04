package types

type CV struct {
	Firstname   string      `yaml:"firstname"`
	Lastname    string      `yaml:"lastname"`
	Job         string      `yaml:"job"`
	Description string      `yaml:"description"`
	Education   Education   `yaml:"education"`
	Experiences Experiences `yaml:"experiences"`
	Skills      []string    `yaml:"skills"`
}

type Education struct {
	Timerange string `yaml:"timerange"`
	Title     string `yaml:"title"`
	Site      string `yaml:"site"`
}

type Experiences struct {
	Timerange   string `yaml:"timerange"`
	Title       string `yaml:"title"`
	Company     string `yaml:"company"`
	Description string `yaml:"description"`
}
