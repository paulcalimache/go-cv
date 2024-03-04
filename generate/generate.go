package generate

import (
	"fmt"
	"os"

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
	fmt.Print(cv.Lastname)
	return nil
}
