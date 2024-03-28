package curriculum

import (
	"bytes"
	"context"
	"log/slog"
	"os"
	"sync"
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

	// Create output directory
	err = os.MkdirAll(output, os.ModePerm)
	if err != nil {
		return err
	}
	err = os.Chdir(output)
	if err != nil {
		return err
	}

	err = saveAsHTML(file)
	if err != nil {
		return err
	}

	err = saveAsPDF(file)
	if err != nil {
		return err
	}

	slog.Info("CV rendered at " + output)
	return nil
}

func saveAsHTML(file bytes.Buffer) error {
	return os.WriteFile("curriculum.html", file.Bytes(), 0644)
}

func saveAsPDF(file bytes.Buffer) error {
	ctx, cancelCtx := chromedp.NewContext(context.Background(), chromedp.WithLogf(slog.Info))
	defer cancelCtx()

	var wg sync.WaitGroup
	wg.Add(1)

	if err := chromedp.Run(ctx,
		chromedp.Navigate("about:blank"),
		chromedp.ActionFunc(func(ctx context.Context) error {
			lctx, cancel := context.WithCancel(ctx)
			chromedp.ListenTarget(lctx, func(ev interface{}) {
				if _, ok := ev.(*page.EventLoadEventFired); ok {
					wg.Done()
					// remove event listener
					cancel()
				}
			})
			return nil
		}),
		chromedp.ActionFunc(func(ctx context.Context) error {
			frameTree, err := page.GetFrameTree().Do(ctx)
			if err != nil {
				return err
			}
			return page.SetDocumentContent(frameTree.Frame.ID, file.String()).Do(ctx)
		}),
		// wait for the page.EventLoadEventFired
		chromedp.ActionFunc(func(ctx context.Context) error {
			wg.Wait()
			return nil
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
				Do(ctx)
			if err != nil {
				slog.Debug("test 5")
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
	stylePath := "templates/" + tmpl + "/" + "style.html"
	return template.Must(template.New(tmplFile).ParseFiles(tmplPath, stylePath))
}
