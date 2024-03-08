package curriculum

import (
	"bytes"
	"context"
	"log/slog"
	"os"
	"text/template"

	"github.com/chromedp/cdproto/page"
	"github.com/chromedp/chromedp"
)

func (cv *CV) Render(output string, tmplName string) error {
	slog.Info("Rendering the " + tmplName + " template ...")
	t := getTemplate(tmplName)

	var file bytes.Buffer

	err := t.ExecuteTemplate(&file, "classic.html", cv)
	if err != nil {
		return err
	}

	os.MkdirAll(output, os.ModePerm)
	os.Chdir(output)

	saveAsHTML(file)
	saveAsPDF(file)

	slog.Info("CV rendered at " + output)
	return nil
}

func saveAsHTML(file bytes.Buffer) error {
	return os.WriteFile("curriculum.html", file.Bytes(), 0644)
}

func saveAsPDF(file bytes.Buffer) error {
	ctx, cancelCtx := chromedp.NewContext(context.Background())
	defer cancelCtx()

	if err := chromedp.Run(ctx,
		chromedp.Navigate("about:blank"),
		chromedp.ActionFunc(func(ctx context.Context) error {
			frameTree, err := page.GetFrameTree().Do(ctx)
			if err != nil {
				return err
			}
			return page.SetDocumentContent(frameTree.Frame.ID, file.String()).Do(ctx)
		}),
		chromedp.ActionFunc(func(ctx context.Context) error {
			buf, _, err := page.
				PrintToPDF().
				WithMarginBottom(0).
				WithMarginTop(0).
				WithMarginRight(0).
				WithMarginLeft(0).
				WithPaperHeight(11.67).
				WithPaperWidth(8.27).
				WithPrintBackground(true).
				WithPreferCSSPageSize(true).
				Do(ctx)
			if err != nil {
				return err
			}
			return os.WriteFile("curriculum.pdf", buf, 0644)
		}),
	); err != nil {
		return err
	}
	return nil
}

func getTemplate(tmpl string) *template.Template {
	tmplFile := tmpl + ".html"
	tmplPath := "templates/" + tmpl + "/" + tmplFile
	return template.Must(template.New(tmplFile).ParseFiles(tmplPath))
}
