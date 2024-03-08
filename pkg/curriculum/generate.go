package curriculum

import (
	"log/slog"
	"os"
	"text/template"
)

func (cv *CV) Render(output string, tmplName string) error {
	slog.Info("Rendering the " + tmplName + " template ...")
	t := getTemplate(tmplName)

	f, err := os.Create(output)
	if err != nil {
		return err
	}
	//html
	err = t.ExecuteTemplate(f, "classic.html", cv)
	if err != nil {
		return err
	}

	//pdf
	// htmlBuf := bytes.NewBufferString("")
	slog.Info("CV rendered at " + output)
	return nil
}

func getTemplate(tmpl string) *template.Template {
	tmplFile := tmpl + ".html"
	tmplPath := "templates/" + tmpl + "/" + tmplFile
	return template.Must(template.New(tmplFile).ParseFiles(tmplPath))
}
