package generate

import (
	"context"
	"os"

	"github.com/paulcalimache/go-cv/templates"
	"github.com/paulcalimache/go-cv/types"
	"gopkg.in/yaml.v3"
)

func Generate(file string, output string, format string) error {
	buf, err := os.ReadFile(file)
	if err != nil {
		return err
	}

	var cv types.CV

	err = yaml.Unmarshal(buf, &cv)
	if err != nil {
		return err
	}

	f, err := os.Create(output)
	if err != nil {
		return err
	}
	templates.Classic(cv).Render(context.Background(), f)

	return nil
}
