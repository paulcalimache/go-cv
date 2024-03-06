package cv

import (
	"bytes"
	"context"
	"log"
	"os"

	"github.com/paulcalimache/go-cv/pkg/pdf"
	"github.com/paulcalimache/go-cv/pkg/types"
	"github.com/paulcalimache/go-cv/templates"
	"gopkg.in/yaml.v3"
)

func Generate(file string, output string) error {
	log.Default().Print("Reading file ... " + file)
	buf, err := os.ReadFile(file)
	if err != nil {
		return err
	}

	var cv types.CV

	err = yaml.Unmarshal(buf, &cv)
	if err != nil {
		return err
	}
	log.Default().Print("File ... " + file)

	f, err := os.Create(output)
	if err != nil {
		return err
	}

	// html
	templates.Classic(cv).Render(context.Background(), f)

	//pdf
	htmlBuf := bytes.NewBufferString("")
	templates.Classic(cv).Render(context.Background(), htmlBuf)
	pdf.ConvertHtmlToPdf(htmlBuf.String())

	return nil
}
